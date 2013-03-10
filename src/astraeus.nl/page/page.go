/**
 * User: rnentjes
 * Date: 2/24/13
 * Time: 2:33 PM
 */
package page

import (
	"net/http"
	"html/template"
)

type Page interface {
	Init (string)
	Post (*http.Request) (string)
	Get (*http.Request) (string)
	Model () (interface {})
	Template () (*template.Template)
}

func HandleRequest(page Page, w http.ResponseWriter, r *http.Request) {
	var redirect string

	if (r.Method == "POST") {
		redirect = page.Post(r)
	} else {
		redirect = page.Get(r)
	}

	if redirect == "" {
		model := page.Model()
		templ := page.Template()

		err := templ.Execute(w, model)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, redirect, http.StatusFound)
	}
}

func MakeHandler(page Page, uri string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page.Init(r.URL.Path[len(uri):])

		HandleRequest(page, w,r)
	}
}
