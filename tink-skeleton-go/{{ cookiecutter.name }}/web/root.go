package web

import "github.com/go-chi/chi"

// RouteBuilder constructs an HTTP Handler for the web package.
type RouteBuilder struct{}

// Routes returns a newly constructed http.Handler for the fasttext microservice.
func (rb *RouteBuilder) Routes() chi.Router {
	r := chi.NewRouter()

	r.Mount("/", (&helloWorld{
		Message: "Hello World\n",
	}).Routes())

	return r
}
