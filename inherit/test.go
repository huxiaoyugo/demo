package main

import (
	"reflect"
	"fmt"
)

type Model interface {
	GetFields()
}

type DbModel struct {

	Name string
	bean func() reflect.Type
}


func (d* DbModel) GetFields() {

	typ := d.GetType().Elem()
	for i := 0; i<typ.NumField(); i++ {
		fmt.Println(typ.Field(i).Name)
	}
}

func (d *DbModel) GetType() reflect.Type {
	return d.bean()
}

func (d *DbModel)getType() reflect.Type {
	return reflect.TypeOf(d)
}



func (d* DbModel) SetValue(val string) {
	d.Name = val
}

func (d DbModel) SetVal(val string) {
	d.Name = val
}

type OneDb struct {
	DbModel

	Name string
	Age int
}

func (o *OneDb) getType() reflect.Type {
	return reflect.TypeOf(o)
}

func NewDbModel() *DbModel {

	db := &DbModel{

	}
	db.bean = db.getType
	return db
}

func NewOne() *OneDb {
	db := &OneDb{}
	db.bean = db.getType
	return db
}


func main() {

	db := NewDbModel()
	db.Name = "h"

	fmt.Println(db.Name)


	db.SetVal("huxiaoyu")

	fmt.Println(db.Name)
}


func GetFields(m Model) {
	m.GetFields()
}