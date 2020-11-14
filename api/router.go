package api

import (
  "fmt"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "net/http"
)

const concurrentRequests = 200

func makeAPIHandler() http.Handler {
  r := chi.NewRouter()

  r.Use(
    middleware.RedirectSlashes,
    middleware.Throttle(concurrentRequests),
    middleware.Recoverer,
  )

  r.Get("/random", foo)
  r.Get("/search", foo)

  return r
}

func foo(w http.ResponseWriter, r *http.Request) {
  if _, err := fmt.Fprintf(w, "Hello world!"); err != nil {
    panic(err)
  }
}
