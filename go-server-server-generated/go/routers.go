/*
 * Wikipedia API
 *
 * This is an api that searches a wikipage
 *
 * API version: 1.0.1
 * Contact: bchristian14@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/VoithAPI/WikipediaAPI/1.0.1/",
		Index,
	},

	Route{
		"WikipagesGet",
		strings.ToUpper("Get"),
		"/VoithAPI/WikipediaAPI/1.0.1/wikipages",
		WikipagesGet,
	},

	Route{
		"WikipagesImageSRCsGet",
		strings.ToUpper("Get"),
		"/VoithAPI/WikipediaAPI/1.0.1/wikipages/ImageSRCs",
		WikipagesImageSRCsGet,
	},

	Route{
		"WikipagesWebsiteDataGet",
		strings.ToUpper("Get"),
		"/VoithAPI/WikipediaAPI/1.0.1/wikipages/WebsiteData",
		WikipagesWebsiteDataGet,
	},

	Route{
		"WikipagesWebsiteLinksGet",
		strings.ToUpper("Get"),
		"/VoithAPI/WikipediaAPI/1.0.1/wikipages/WebsiteLinks",
		WikipagesWebsiteLinksGet,
	},

	Route{
		"WikipagesWebsitePagesGet",
		strings.ToUpper("Get"),
		"/VoithAPI/WikipediaAPI/1.0.1/wikipages/WebsitePages",
		WikipagesWebsitePagesGet,
	},
}
