package batchUpdateEng

import (
	"github.com/go-xorm/xorm"
	"reflect"
	"fmt"
	"strings"
	"strconv"
	"errors"
	"github.com/qianlnk/log"
)

// 一次update最多的对象个数
const DefaultMaxUpdateCount = 200

type batchUpdateEng struct {
	eng        OrmEng      // 数据库引擎
	bOpt       BatchOption // 配置对象
	cols       []string    // 待更新的列
	beansBatch BeanBatch   // 分批处理待更新的beans
	pk         string      // 主键名称
	tableName  string      // 表明
}

func NewBatchUpdateEngine(eng *xorm.Engine, opts ...Option) *batchUpdateEng {
	bue := &batchUpdateEng{
		eng:        NewXormEngine(eng),
		beansBatch: &beanBatch{},
		bOpt:       apply(&opts),
	}
	if bue.bOpt.onceMaxCount == 0 {
		bue.bOpt.onceMaxCount = DefaultMaxUpdateCount
	}
	return bue
}

func (eng *batchUpdateEng) Cols(cols ... string) *batchUpdateEng {
	eng.cols = eng.cols[0:0]
	eng.cols = append(eng.cols, cols...)
	return eng
}

func (eng *batchUpdateEng) Table(tableName string) *batchUpdateEng {
	eng.tableName = tableName
	return eng
}

func (eng *batchUpdateEng) Pk(primaryKey string) *batchUpdateEng {
	eng.pk = primaryKey
	return eng
}

func (eng *batchUpdateEng) Update(beans ... interface{}) (affected int64, err error) {
	defer eng.clear()

	eng.beansBatch.Init(beans)
	for {
		beanArr, has := eng.beansBatch.Next(eng.bOpt.onceMaxCount)
		if !has {
			break
		}
		var af int64
		af, err = eng.update(beanArr)
		if err != nil {
			return affected, err
		}
		affected += af
	}
	return affected, nil
}

func (eng *batchUpdateEng) update(beans []interface{}) (int64, error) {
	// 转换为批量update的sql语句
	sql, err := eng.createdBatchUpdateSQL(beans)
	if err != nil {
		return 0, err
	}

	result, err := eng.eng.Exec(sql)
	if err != nil {
		if eng.isShowLog() {
			log.Error("update er: %v", err)
		}
	}
	return result, err
}

func (eng *batchUpdateEng) createdBatchUpdateSQL(beans []interface{}) (string, error) {
	var (
		ids string
	)
	valMaps, err := eng.getValMaps(beans)

	if err != nil {
		return "", nil
	}
	// Ids
	for k, _ := range valMaps {
		ids += fmt.Sprintf("%d,", k)
	}
	// 去掉最后一个逗号
	ids = ids[0 : len(ids)-1]

	preSql := fmt.Sprintf("update %s set ", eng.tableName)

	colsql := ""
	for index, col := range eng.cols {
		sql := col + " = case " + eng.pk
		for k, val := range valMaps {
			sql += fmt.Sprintf(" when %d then '%v' ", k, (*val)[col])
		}
		sql += " end "
		if index != len(eng.cols)-1 {
			sql += ", "
		}
		colsql += sql
	}

	whereSql := fmt.Sprintf(" where %s in (%s)", eng.pk, ids)

	resultSql := preSql + colsql + whereSql

	if eng.isShowLog() {
		log.Info(resultSql)
	}

	return resultSql, nil
}

func (eng *batchUpdateEng) getValMaps(beans []interface{}) (map[int]*map[string]interface{}, error) {
	valMaps := make(map[int]*map[string]interface{})
	var err error
	if len(beans) == 0 {
		return valMaps, errors.New("待更新的模型个数为0")
	}

	if eng.pk == "" {
		eng.pk, err = eng.eng.GetPrimaryKey(beans[0])
		if err != nil {
			return valMaps, err
		}
	}
	if eng.tableName == "" {
		eng.tableName = eng.eng.GetTableName(beans[0])
	}

	if len(eng.cols) == 0 {
		columns := eng.eng.Columns(beans[0])
		eng.cols = append(eng.cols, columns...)
	}

	for _, bean := range beans {
		sliceValue := reflect.Indirect(reflect.ValueOf(bean))
		if sliceValue.Kind() != reflect.Struct {
			return valMaps, errors.New("update bean is not struct")
		}
		pkVal, err := eng.getPrimaryKeyVal(bean)
		if err != nil {
			return valMaps, err
		}
		vMap := make(map[string]interface{}, 0)
		valMaps[pkVal] = &vMap

		// 获取每个字段的值
		for _, col := range eng.cols {
			vMap[col] = eng.getValue(bean, col)
		}
	}
	return valMaps, nil
}

func (eng *batchUpdateEng) getPrimaryKeyVal(bean interface{}) (val int, err error) {

	if eng.pk == "" {
		eng.pk, err = eng.eng.GetPrimaryKey(bean)
		if err != nil {
			return 0, err
		}
	}
	if err != nil && eng.isShowLog() {
		log.Error("getPrimaryKeyVal err:%v", err)
	}
	return interfaceToInt(eng.getValue(bean, eng.pk))
}

func (eng *batchUpdateEng) getValue(bean interface{}, key string) interface{} {

	if key == "" {
		if eng.isShowLog() {
			log.Warn("GetValue key is empty")
		}
		return nil
	}

	beanValue := reflect.ValueOf(bean)

	val := reflect.Indirect(beanValue)

	key = toCamelCase(key)
	return val.FieldByName(key).Interface()
}

func (eng *batchUpdateEng) isShowLog() bool {
	return eng.bOpt.isShowLog
}

func (eng *batchUpdateEng) clear() {
	eng.pk = ""
	eng.tableName = ""
	eng.beansBatch.Clear()
	eng.cols = make([]string, 0)
}

func toCamelCase(key string) string {

	res := ""
	arr := strings.Split(key, "_")
	for _, item := range arr {
		for index, char := range item {
			if index == 0 && char >= 97 && char <= 122 {
				res += string(char - 32)
			} else {
				res += string(char)
			}
		}
	}
	return res
}

func interfaceToInt(bean interface{}) (int, error) {

	switch reflect.TypeOf(bean).Kind() {
	case reflect.Int, reflect.Int64, reflect.Int16, reflect.Int32, reflect.Int8,
		reflect.Uint, reflect.Uint64, reflect.Uint16, reflect.Uint32, reflect.Uint8:
		val, err := strconv.Atoi(fmt.Sprintf("%v", bean))
		return val, err
	default:
		return 0, errors.New("目前支持主键为整形或者字符的")
	}
}
