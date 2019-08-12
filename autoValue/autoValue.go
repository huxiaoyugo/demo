package modelAutoValue

import (
	"reflect"
	"fmt"
	"errors"
)

type FieldModel struct {
	Name   string
	Ignore  bool
	ValFunc func()(interface{}, error)
}


type CanSetCol interface {
	SetField(fieldName string, val interface{}) error
}

type InnerField interface {
	InnerFieldValue()
}

type AutoValueInterface interface {
	NewModel() interface{}
}

type ModelAutoValue struct {
	fields []*FieldModel
	Bean AutoValueInterface
}

func NewModelAutoValue(bean AutoValueInterface) (*ModelAutoValue, error){
	m := &ModelAutoValue{
		Bean:bean,
	}
	// 先检查传入的bean是否合法
	err :=m.checkBean()
	if err != nil {
		return nil, err
	}

	// 默认设置所有的项为需要计算的项
	m.setAllIgnoreFields()
	return m,err
}

func (this* ModelAutoValue) Fields(models... *FieldModel) * ModelAutoValue {
	this.fields = this.fields[0:0]
	this.fields = append(this.fields, models...)
	return this
}

func (this* ModelAutoValue) Ignore(models... string) * ModelAutoValue {
	this.setAllIgnoreFields(models...)
	return this
}


func (this* ModelAutoValue)setAllIgnoreFields(models... string) {
	m := this.Bean.NewModel()

	val := reflect.TypeOf(m).Elem()

	fieldArr := make([]*FieldModel, 0)
out: for i:=0; i<val.NumField(); i++ {
	fieldName :=val.Field(i).Name
	for _, item := range models {
		if item == fieldName {
			continue out
		}
	}
	fieldArr = append(fieldArr, &FieldModel{Name:fieldName,Ignore:false})
}
	this.fields = fieldArr
}



func (this* ModelAutoValue)checkBean() error {
	bean := this.Bean
	if bean == nil {
		return errors.New("为初始化Bean")
	}

	if reflect.TypeOf(bean).Kind() != reflect.Ptr {
		return  errors.New("bean不是指针")
	}

	m := bean.NewModel()
	if reflect.TypeOf(m).Kind() != reflect.Ptr {
		return errors.New("bean.NewModel()不是指针")
	}

	val := reflect.TypeOf(m).Elem()

	if val.Kind() != reflect.Struct {
		return errors.New("bean.NewModel()不是结构体指针")
	}
	return nil
}

func (this * ModelAutoValue) checkFunc() error {
	bean := this.Bean
	m := bean.NewModel()
	val := reflect.TypeOf(m).Elem()

	count := val.NumField()
	fieldArr := make([]string,count)
	for i:=0; i<count; i++ {
		fieldArr[i] = val.Field(i).Name
	}

out1: for _, item := range this.fields {
	for _, itemVal := range fieldArr {
		if item.Name == itemVal {
			continue out1
		}
	}
	return errors.New(fmt.Sprintf("%v模型中不存在%s字段", val.Name(), item.Name))
}

	val = reflect.TypeOf(bean).Elem()
	count = val.NumMethod()
	funcArr := make([]string,count)
	for i:=0; i<count;i++ {
		funcArr[i] = val.Method(i).Name
	}

out2:for _,item := range this.fields {
	if item.Ignore {
		continue
	}
	for _,itemVal := range funcArr {
		if itemVal == item.Name {
			continue out2
		}
	}
	return errors.New(fmt.Sprintf("%s字段没有对应的方法", item.Name))
}
	return nil
}

func (this* ModelAutoValue) Get() (interface{},error) {

	// 检查需要计算的字段是否有对应的函数
	err := this.checkFunc()
	if err != nil {
		return nil,err
	}
	bean := this.Bean
	resModel := bean.NewModel()

	val := reflect.ValueOf(bean).Elem()

	for _,item := range this.fields {

		if item.Ignore {continue}

		resVal := reflect.ValueOf(resModel).Elem().FieldByName(item.Name)
		if !resVal.CanSet() {
			return resModel, errors.New(fmt.Sprintf("%v属性不能被设置", item.Name))
		}

		var funcVal interface{}

		inputs := make([]reflect.Value, 0)

		returnValues := val.MethodByName(item.Name).Call(inputs)

		if len(returnValues) != 2 {
			return resModel, errors.New("函数执行的返回值不是2个")
		}

		if !returnValues[1].IsNil() {
			return resModel, errors.New(fmt.Sprintf("func %v() err:%v", item.Name, returnValues[1]))
		}
		funcVal = returnValues[0].Interface()


		valKind := reflect.TypeOf(funcVal).Kind()
		switch valKind {
		case reflect.Int:
			resVal.SetInt(int64(funcVal.(int)))
		case reflect.Int8:
			resVal.SetInt(int64(funcVal.(int8)))
		case reflect.Int16:
			resVal.SetInt(int64(funcVal.(int16)))
		case reflect.Int32:
			resVal.SetInt(int64(funcVal.(int32)))
		case reflect.Int64:
			resVal.SetInt(int64(funcVal.(int64)))
		case reflect.String:
			resVal.SetString(funcVal.(string))
		default:
			return nil, errors.New("函数返回值和字段类型无法匹配")
		}
	}

	resModel.InnerFieldValue()
	return resModel,nil
}



func (this* ModelAutoValue) GetValues() (error) {

	// 检查需要计算的字段是否有对应的函数
	err := this.checkFunc()
	if err != nil {
		return err
	}
	bean := this.Bean
	//resModel := bean.NewModel()

	val := reflect.ValueOf(bean).Elem()
	inputs := make([]reflect.Value, 0)

	for _,item := range this.fields {

		if item.Ignore {continue}

		returnValues := val.MethodByName(item.Name).Call(inputs)

		if len(returnValues) != 2 {
			return errors.New("函数执行的返回值不是2个")
		}

		if !returnValues[1].IsNil() {
			return errors.New(fmt.Sprintf("func %v() err:%v", item.Name, returnValues[1]))
		}
	}
	return nil
}
