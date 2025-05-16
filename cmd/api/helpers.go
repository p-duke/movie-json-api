package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Retrieve the "id" URL parameter from the current request context, then convert it to
// an integer and return it. If the operation isn't successful, return 0 and an error.
// The readIDParam() method doesn’t use any dependencies from our application struct so it could just be a regular function, rather than a method on application. But in general, I suggest setting up all your application-specific handlers and helpers so that they are methods on application. It helps maintain consistency in your code structure, and also future-proofs your code for when those handlers and helpers change later and they do need access to a dependency.
func (app *application) readIDParam(r *http.Request) (int64, error) {
	// When httprouter is parsing a request, any interpolated URL parameters will be
	// stored in the request context. We can use the ParamsFromContext() function to
	// retrieve a slice containing these parameter names and values.
	params := httprouter.ParamsFromContext(r.Context())

	// We can then use the ByName() method to get the value of the "id" parameter from
	// the slice. In our project all movies will have a unique positive integer ID, but
	// the value returned by ByName() is always a string. So we try to convert it to a
	// base 10 integer (with a bit size of 64). If the parameter couldn't be converted,
	// or is less than 1, we know the ID is invalid so we use the http.NotFound()
	// function to return a 404 Not Found response.
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data any, headers http.Header) error {
	// Encode the data to JSON, returning error if there was one
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a new line to improve viewing in terminal applications
	js = append(js, '\n')

	// At this point, we know that we won't encounter any more errors before writing the
    // response, so it's safe to add any headers that we want to include. We loop
    // through the header map and add each header to the http.ResponseWriter's header map.
    // Note that it's OK if the provided header map is nil. Go doesn't throw an error
    // if you try to range over (or generally, read from) a nil map.
	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
