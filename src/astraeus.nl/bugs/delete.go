/**
 * User: rnentjes
 * Date: 3/9/13
 * Time: 11:09 PM
 */
package bugs

import (
	"net/http"
	"html/template"
	"strconv"
)

var deleteTemplate = template.Must(template.ParseFiles("data/bugs/html/delete.html"))

type BugsDeletePage struct {
	bug *Bug
}

func (mp *BugsDeletePage) Init(uri string) {
	id, _ := strconv.ParseUint(uri, 10, 64)

	mp.bug = GetBugs().GetBug(id)
}

func (mp *BugsDeletePage) Get(r *http.Request) string {
	return ""
}

func (mp *BugsDeletePage) Post(r *http.Request) string {
	if r.FormValue("action") == "Delete" {
		GetBugs().DeleteBug(mp.bug)
	}

	return "/overview"
}

func (mp *BugsDeletePage) Model() interface {} {
	return mp.bug
}

func (mp *BugsDeletePage) Template() *template.Template {
	return deleteTemplate
}
