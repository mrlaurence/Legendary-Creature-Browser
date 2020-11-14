package api

import (
  "encoding/json"
  "fmt"
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
      w.WriteHeader(http.StatusBadRequest)
      return
    }
    f(w, r, c, n, vs)
  }
}

func randomAPI(w http.ResponseWriter, r *http.Request, c creatures, n int, vs url.Values) {
  if _, err := fmt.Fprintf(w, "do i look like i'm implemented yet"); err != nil {
    panic(err)
  }
}

func searchAPI(w http.ResponseWriter, r *http.Request, c creatures, n int, vs url.Values) {
  q := vs.Get("q")
  sStr := vs.Get("s")
  sPcnt, err := strconv.Atoi(sStr)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  s := float64(sPcnt) / 100.0
  if s < 0 {
    s = 0
  } else if s > 1 {
    s = 1
  }
  res := c.search(q, n, s).toModel()
  if err := json.NewEncoder(w).Encode(res); err != nil {
    panic(err)
  }
}
