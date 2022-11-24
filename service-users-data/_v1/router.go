// Package api regroup all http related files
package api

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// InitRouter function set up router & middlewares and mount different routes
func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),

		middleware.RedirectSlashes,
		middleware.Logger,
		middleware.Recoverer,

		middleware.Heartbeat("/h"),

		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)

	r.Use(middleware.Timeout(300 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/users", UserRoutes())
		// r.Mount("/metrics", nil) // for futur implementation of monitoring agent
	})

	return r
}
