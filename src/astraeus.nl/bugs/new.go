/**
 * User: rnentjes
 * Date: 3/9/13
 * Time: 11:09 PM
 */
package bugs

import (
	"net/http"
	"html/template"
)

var newTemplate = template.Must(template.ParseFiles("data/bugs/html/new.html"))

type BugsNewPage struct {
}

func (mp *BugsNewPage) Init(uri string) {
}

func (mp *BugsNewPage) Get(r *http.Request) string {
	return ""
}

func (mp *BugsNewPage) Post(r *http.Request) string {
	if r.FormValue("action") == "Save" {
		bug := CreateBug(r.FormValue("title"), r.FormValue("body"))

		GetBugs().SaveBug(bug)
	}

	return "/overview"
}

func (mp *BugsNewPage) Model() interface {} {
	var bug Bug

	return bug
}

func (mp *BugsNewPage) Template() *template.Template {
	return newTemplate
}
