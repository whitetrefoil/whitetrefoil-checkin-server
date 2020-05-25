package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"whitetrefoil.com/checkin/server/jr"
)

type Config struct {
	AppId     string
	AppSecret string
	Redirect  string
}

func NewServer(port int, id string, sec string, red string) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.StripSlashes)

	r.With(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "cfg", &Config{
				AppId:     id,
				AppSecret: sec,
				Redirect:  red,
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}).Route("/api", func(r chi.Router) {
		r.Get("/login", getLoginUrl)
		r.Post("/login", checkLogin)

		r.With(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				token := r.Header.Get("X-Token")
				if token == "" {
					jr.Json401(w, nil)
					return
				}
				ctx := context.WithValue(r.Context(), "token", token)
				next.ServeHTTP(w, r.WithContext(ctx))
			})
		}).Group(func(r chi.Router) {
			r.Get("/users", getUserDetail)
			r.Get("/venues", searchVenues)
			r.Post("/checkin", addCheckin)
		})
	})

	log.Printf("Started on :%d\n", port)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}
}
