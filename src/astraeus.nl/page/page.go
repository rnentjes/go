/**
 * User: rnentjes
 * Date: 2/24/13
 * Time: 2:33 PM
 */
package page

import (
	"io"
	"net/http"
)

type Renderer interface {
	Render(w *io.Writer)
}

type Page interface {

	Get(parameters map[string]([]string))
	Post()
	ServeHttp(w http.ResponseWriter, r *http.Request)

}

type PageMapping struct {
	pages map[string]*Page
}

func (pm *PageMapping) Add(uri string, page *Page) {
	pm.pages[uri] = page
}

func (pm *PageMapping) Get(uri string) *Page {
	return pm.pages[uri]
}




