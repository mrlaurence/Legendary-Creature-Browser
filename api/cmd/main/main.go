package main

import (
  "context"
  "github.com/mrlaurence/Tech-Support-Gurus-Oxford-Hack-2020/api"
  "math/rand"
  "os"
  "os/signal"
  "syscall"
  "time"
)

func main() {
  rand.Seed(time.Now().Unix())

  interrupt := make(chan os.Signal, 1)
  signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
  defer signal.Stop(interrupt)

  err, fatal, shutdown := api.Serve("config.json")
  if err != nil {
    panic(err)
  }

  select {
  case <-interrupt:
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()
    if err := shutdown(ctx); err != nil {
      panic(err)
    }
  case err = <-fatal:
    panic(err)
  }
}
