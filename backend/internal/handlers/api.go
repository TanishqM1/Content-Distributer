package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// in this file, I need to setup the handler. While it typically uses middleware, and we re-route to that here, we don't have any middleware! there is no permissions-based
// things that we need to handle here.

func Handler(r *chi.Mux) {
	// strip trailing slashes (from chi package)
	r.Use(chimiddle.StripSlashes)

	// setup route (no middleware from our end)
	// Need to implement the PostContent function still!

	r.Route("/post", func(router chi.Router) {
		// implementation for this endpoint
		// ==> router.Get("/content", PostContent)
	})
}
