package web

import (
	"html/template"
	"net/http"

	"github.com/connorjcantrell/syndicate"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHandler(store syndicate.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
	h.Use(middleware.Logger)
	h.Route("/manufacturers", func(r chi.Router) {
		r.Get("/", h.ManufacturerList())
	})

	return h
}

type Handler struct {
	*chi.Mux

	store syndicate.Store
}

const ManuListHTML = `
<h1>Manufacturers</h1>
{{range .Manufacturers}}
	<dt><strong>{{.Name}}</strong></dt>
	<dd>{{.Active}}</dd>
{{end}}
</dt>
`

func (h *Handler) ManufacturerList() http.HandlerFunc {
	type data struct {
		Manufacturers []syndicate.Manufacturer
	}
	tmpl := template.Must(template.New("").Parse(`ManuListHTML`))
	return func(w http.ResponseWriter, r *http.Request) {
		mm, err := h.store.Manufacturers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data{Manufacturers: mm})
	}
}
