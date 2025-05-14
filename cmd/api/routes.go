package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	// Register the relevant methods, URL patterns and handler functions for our
    // endpoints using the HandlerFunc() method. Note that http.MethodGet and
    // http.MethodPost are constants which equate to the strings "GET" and "POST"
    // respectively.
	// There are a couple of benefits to encapsulating our routing rules in this way. The first benefit is that it keeps our main() function clean and ensures all our routes are defined in a single place. The other big benefit, which we demonstrated in the first Let’s Go book, is that we can now easily access the router in any test code by initializing an application instance and calling the routes() method on it
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	// Return the httprouter instance
	return router
}
