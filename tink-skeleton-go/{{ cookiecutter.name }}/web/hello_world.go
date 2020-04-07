package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tink-ab/tink-go-libraries/web"
)

type helloWorld struct {
	Message string
}

// Routes returns an HTTP handler that allows importing and exporting a model.
func (hw *helloWorld) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.With(web.InstrumentationMiddleware("/")).Get("/", hw.reply)
	})
	return r
}

func (hw *helloWorld) reply(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(hw.Message))
}
