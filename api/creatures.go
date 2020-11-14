package api

import (
  "encoding/json"
  "io/ioutil"
)

type creatures struct {
  Name string      `json:"Name"`
  Desc description `json:"Description"`
  Img  string      `json:"Img"`
}

type description []string

func readCreatures(path string) (creatures, error) {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    return creatures{}, err
  }
  var c creatures
  if err := json.Unmarshal(data, &c); err != nil {
    return creatures{}, err
  }
  return c, nil
}
