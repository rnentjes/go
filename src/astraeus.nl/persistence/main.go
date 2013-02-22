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
	persistence.PersistentParent
	name 	string
}

type OtherData struct {
	persistence.PersistentParent
	Firstname 	string
	Lastname 	string
	Child		*persistence.Persistent
}

var store *persistence.PersistentStore

func main() {
	store = persistence.Open("path")

	var md *MyData
	var od *OtherData
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

	fmt.Println(reflect.ValueOf(od))
	fmt.Println(reflect.ValueOf(od))

	fmt.Println(store)
	fmt.Println(store.Data["*persistence.PersistentParent"])
	fmt.Println(store.Data["*persistence.PersistentParent"][1])

	var pers = store.Find("*persistence.PersistentParent", 2)

	fmt.Println(pers)

	var other = store.Find("*persistence.PersistentParent", 3)

	fmt.Println(other)
}
