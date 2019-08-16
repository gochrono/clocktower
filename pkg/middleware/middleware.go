package middleware

import (
	"github.com/NYTimes/gziphandler"
	"github.com/gochrono/castle/config"
	"github.com/justinas/alice"
	_ "github.com/lib/pq" // Load postges drivers
	"github.com/rs/cors"
	"net/http"
)

// applyGlobalMiddleware applies the common middleware to the whole app
func ApplyGlobalMiddleware(app http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins: config.CorsAllowedOrigins(),
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Authorization", "Accept", "Content-Type"},
	})

	return alice.New(
		cors.Handler,
		gziphandler.GzipHandler,
	).Then(app)
}
