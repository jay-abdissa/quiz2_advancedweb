//Filename:cmd/api/routes.go
package main

import	(
	"net/http"

	"github.com/julienschmidt/httprouter"
)
func (app *application) routes() *httprouter.Router{
	//Create new httprouter instance
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed =http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/entries", app.createEntryHandler)
	router.HandlerFunc(http.MethodGet, "/v1/rand/:id", app.showRandomHandler)
	return router
}

