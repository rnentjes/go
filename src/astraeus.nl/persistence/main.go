/**
 * User: rnentjes
 * Date: 2/22/13
 * Time: 8:46 PM
 */
package main

import (
	"fmt"
	"astraeus.nl/persistence"
	"reflect"
)

type MyData struct {
	persistence.PersistentStruct
	name 	string
}

type OtherData struct {
	persistence.PersistentStruct
	Firstname 	string
	Lastname 	string
	Child		*persistence.Persistent
}

func (td *OtherData) Type() string {
	return "OtherData"
}

type TestData struct {
	id			uint64
	Firstname 	string
	Lastname 	string
}

func (td *TestData) Type() string {
	return "TestData"
}

func (td *TestData) Id() uint64 {
	return td.id
}

func (td *TestData) GenerateId() {
	td.id = persistence.NextId()
}

func (td *TestData) Clone() persistence.Persistent {
	var result *TestData = new (TestData)

	result.id 			= td.id
	result.Firstname 	= td.Firstname
	result.Lastname 	= td.Lastname

	return result
}

var store *persistence.PersistentStore

func main() {
	store = persistence.Open("path")

	var md *MyData
	var od *OtherData
	var td *TestData

	md = new(MyData);
	md.name = "Rien"

	store.Save(md)

	md = new(MyData);
	md.name = "Pipo"

	store.Save(md)

	od = new(OtherData)
	od.Firstname	= "Rien"
	od.Lastname 	= "Nentjes"

	store.Save(od)

	td = new(TestData)
	td.Firstname	= "Rien2"
	td.Lastname 	= "Nentjes2"

	store.Save(td)

	fmt.Println(reflect.ValueOf(od))
	fmt.Println(reflect.ValueOf(od))

	fmt.Println(store)
	fmt.Println(store.Data["*persistence.PersistentStruct"])
	fmt.Println(store.Data["*persistence.PersistentStruct"][1])

	var pers = store.Find("*persistence.PersistentStruct", 2)

	fmt.Println(pers)

	var other = store.Find("*persistence.PersistentStruct", 3)

	fmt.Println(other)

	other = store.Find("TestData", 4)

	fmt.Println(other)
}
