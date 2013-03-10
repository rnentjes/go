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
	"strconv"
	"astraeus.nl/page"
)

var overviewTemplate = template.Must(template.ParseFiles("data/bugs/html/notfound.html"))

func handleNewBug(w http.ResponseWriter, r *http.Request) {
	bug1 := bugs.CreateBug("")

	bugs.GetBugs().SaveBug(bug1)

	http.Redirect(w, r, "/edit/"+strconv.FormatUint(bug1.Id(), 10), http.StatusFound)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := overviewTemplate.ExecuteTemplate(w, "notfound.html", bugs.GetBugs())

	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/notfound.html", notFoundHandler)
	http.HandleFunc("/overview", page.MakeHandler(new(bugs.BugsOverviewPage), "/overview"))
	http.HandleFunc("/edit/", page.MakeHandler(new(bugs.BugsEditPage), "/edit/"))
	http.HandleFunc("/new", handleNewBug)

	http.ListenAndServe(":8080", nil)
}
