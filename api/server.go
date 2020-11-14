package api

import (
  "context"
  "fmt"
  "net/http"
  "time"
)

type config struct {
  Port uint `json:"port"`
}

type ShutdownFunc func(context.Context) error

func Serve(configPath string) (error, <-chan error, ShutdownFunc) {
  c, err := readConfig(configPath)
  if err != nil {
    return err, nil, nil
  }

  server := http.Server{
    Addr:        fmt.Sprintf(":%d", c.Port),
    ReadTimeout: time.Second * 10,
    Handler:     makeAPIHandler(),
  }

  fatal := make(chan error, 1)

  go func() {
    if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      fatal <- err
    }
  }()

  return nil, fatal, func(ctx context.Context) error {
    return server.Shutdown(ctx)
  }
}

func readConfig(path string) (config, error) {
  var c config
  if err := readJSONFile(path, &c); err != nil {
    return config{}, err
  }
  return c, nil
}
