/**
 * User: rnentjes
 * Date: 2/26/13
 * Time: 10:19 PM
 */
package main

import (
	"net/http"
	"html/template"
	"astraeus.nl/bugs"
	"fmt"
)

var overviewTemplate = template.Must(template.ParseFiles("data/bugs/html/overview.html"))
// template.New("overview")
// templates = template.Must(template.ParseFiles("data/bugs/html/overview.html"))

func handleOverview(w http.ResponseWriter, r *http.Request) {
	bug1 := bugs.CreateBug("Test")
	bug2 := bugs.CreateBug("Problem?")

	myBugs.AddBug(bug1)
	myBugs.AddBug(bug2)

	err := overviewTemplate.ExecuteTemplate(w, "overview.html", myBugs.Bugs)
	if err != nil {
		panic(err)
	}
}

var myBugs = bugs.CreateBugs()

func main() {
	/*
	var err error
	overviewTemplate, err = overviewTemplate.ParseFiles("data/bugs/html/overview.html")
	if err != nil {
		panic(err)
	}*/
	fmt.Println(overviewTemplate.Name())
	http.HandleFunc("/overview", handleOverview)
	http.ListenAndServe(":8080", nil)
}
