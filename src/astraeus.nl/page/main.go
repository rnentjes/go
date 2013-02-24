/**
 * User: rnentjes
 * Date: 2/24/13
 * Time: 3:42 PM
 */
package page

type MainPage struct {
	user string
}

func (mp *MainPage) Init() {
	// create template
}

func (mp *MainPage) Get(parameters map[string]([]string)) {
	// get user from string
}

/*
	page := pageMapping.get(currentUri)

	page.Render(session, parameters, writer, request);
*/
