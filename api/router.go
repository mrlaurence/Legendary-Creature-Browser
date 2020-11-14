package api

import (
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "net/http"
  "net/url"
  "strconv"
)

const concurrentRequests = 200

type creaturesAPIFunc func(w http.ResponseWriter, r *http.Request, c creatures, n int, vs url.Values)

func makeAPIHandler(conf config) http.Handler {
  r := chi.NewRouter()

  r.Use(
    middleware.RedirectSlashes,
    middleware.Throttle(concurrentRequests),
    middleware.Recoverer,
  )

  r.Get("/random", mwCreatures(conf.CreaturesPath, randomAPI))
  r.Get("/search", mwCreatures(conf.CreaturesPath, searchAPI))

  return r
}

func mwCreatures(path string, f creaturesAPIFunc) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    c, err := readCreatures(path)
    if err != nil {
      panic(err)
    }
    vs := r.URL.Query()
    nStr := vs.Get("n")
    n, err := strconv.Atoi(nStr)
    if err != nil {
      panic(err)
    }
    f(w, r, c, n, vs)
  }
}

func randomAPI(w http.ResponseWriter, r *http.Request, c creatures, n int, vs url.Values) {

}

func searchAPI(w http.ResponseWriter, r *http.Request, c creatures, n int, vs url.Values) {

}
