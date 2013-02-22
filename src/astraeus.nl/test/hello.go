/**
 * User: rnentjes
 * Date: 2/22/13
 * Time: 8:04 PM
 */
package main

import (
	"fmt"
	"net/http"
)

type Page struct {

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, url: %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	fmt.Printf("Started.\n")
}
