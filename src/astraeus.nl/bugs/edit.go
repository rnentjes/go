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

var editTemplate = template.Must(template.ParseFiles("data/bugs/html/edit.html"))

type BugsEditPage struct {
	bug *Bug
}

func (mp *BugsEditPage) Init(uri string) {
	id, _ := strconv.ParseUint(uri, 10, 64)

	mp.bug = GetBugs().GetBug(id)
}

func (mp *BugsEditPage) Get(r *http.Request) string {
	return ""
}

func (mp *BugsEditPage) Post(r *http.Request) string {
	if r.FormValue("action") == "Save" {
		mp.bug.Description = r.FormValue("body")
		mp.bug.Title = r.FormValue("title")
		GetBugs().SaveBug(mp.bug)
	}

	return "/overview"
}

func (mp *BugsEditPage) Model() interface {} {
	return mp.bug
}

func (mp *BugsEditPage) Template() *template.Template {
	return editTemplate
}
