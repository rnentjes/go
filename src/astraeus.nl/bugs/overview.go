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

var overviewTemplate = template.Must(template.ParseFiles("data/bugs/html/overview.html"))

/*
	Init(uri string)
	Post(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Model() interface {}
	Template() *template.Template
 */

type BugsOverviewPage struct {
	user string
}

func (mp *BugsOverviewPage) Init(uri string) {
}

func (mp *BugsOverviewPage) Get(r *http.Request) string {
	return ""
}

func (mp *BugsOverviewPage) Post(r *http.Request) string {
	return ""
}

func (mp *BugsOverviewPage) Model() interface {} {
	return GetBugs().Bugs
}

func (mp *BugsOverviewPage) Template() *template.Template {
	return overviewTemplate
}
