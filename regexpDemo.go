package main

import (
	"strings"
	"regexp"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	str := "你好啊"


	by := *(*[]byte)(unsafe.Pointer(&str))

	for _,item := range by {
		fmt.Printf(" %d", item)
	}

	fmt.Println()
	for _, item := range str {
		fmt.Printf("%s", string(item))
	}
	fmt.Println()

	r := []rune(str)

	fmt.Println(len(r))
}

func (f* Father) Func2() {
	fmt.Println("Father Func1")
}

func (f *Father) func2() {
	fmt.Println("Father Func2")

	fmt.Println("type:",reflect.TypeOf(f).Elem().Name())
}


type Father struct {
	Tax  string `json:"tax"`
}


func (f* Mode) Func1(a int ) {

	fmt.Println("Model Func1")
}

type Mode struct {
	Father
}

func  GetValue(bean interface{}, key string ) interface{} {

	beanValue := reflect.ValueOf(bean)

	val := beanValue
	if beanValue.Type().Kind() == reflect.Ptr {
		val = beanValue.Elem()
	}

	return val.FieldByName(toCamelCase(key)).Interface()
}


func toCamelCase(key string) string{

	res := ""
	arr := strings.Split(key, "_")
	for _, item := range arr {
		for index , char := range item {
			if index == 0 && char>=97 && char<=122 {
				res += string(char-32)
			} else {
				res += string(char)
			}
		}
	}
	return res
}


func delete_extra_space(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)      //替换tab为空格
	regstr := "\\s{2,}"                          //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(regstr)             //编译正则表达式
	s2 := make([]byte, len(s1))                  //定义字符数组切片
	copy(s2, s1)                                 //将字符串复制到切片
	spc_index := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spc_index) > 0 {                     //找到适配项
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...) //删除多余空格
		spc_index = reg.FindStringIndex(string(s2))            //继续在字符串中搜索
	}
	return string(s2)
}


func ignore(sql string) string {

	resultArr := make([]string, 0)
	tempArr := make([]string, 0)

	index := -1
	strArr := strings.Split(sql," ")

	for i, item := range strArr {

		if strings.Contains(item,"?") {
			index = i
			break
		}
		//fmt.Println(item)
	}

	for i := index; i >= 0; i-- {
		if hasWhereAndOrGroupOrder(strArr[i]) {
			break
		}
		tempArr = append(tempArr, strArr[i])
	}


	for i:= len(tempArr)-1; i>=0 ;i-- {
		resultArr = append(resultArr, tempArr[i])
	}

	tempArr = tempArr[0:0]
	for i := index+1; i < len(strArr); i++ {
		if hasWhereAndOrGroupOrder(strArr[i]) {
			break
		}
		tempArr = append(tempArr, strArr[i])
	}

	resultArr = append(resultArr,tempArr...)

	resultStr := strings.Join(resultArr," ")

	fmt.Println(resultStr)
	// 如果原来有阔括号，需要保留括号
	resultStr = strings.Replace(resultStr,"(","",-1)
	resultStr = strings.Replace(resultStr,")","",-1)

	sql = strings.Replace(sql,resultStr," 1=1 ",1)
	return sql
}

func hasWhereAndOrGroupOrder(str string) bool {

	reg, err := regexp.Compile(`\(?((and)|(where)|(or)|(group)|(order))$`)

	if err != nil {
		fmt.Println("hasWhereAndOr:", err)
		return false
	}

	return reg.MatchString(str)
}