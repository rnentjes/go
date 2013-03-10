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
	"astraeus.nl/page"
)

var overviewTemplate = template.Must(template.ParseFiles("data/bugs/html/notfound.html"))

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
	http.HandleFunc("/new", page.MakeHandler(new(bugs.BugsNewPage), ""))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/overview", http.StatusFound)
		})

	http.ListenAndServe(":7000", nil)
}
