package middleware

import (
	"strings"

	"github.com/go-chi/cors"
)

func CORS(allowedOrigins string) cors.Options {
	origins := []string{"*"}
	if allowedOrigins != "" && allowedOrigins != "*" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
		}
	}

	return cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}
}
