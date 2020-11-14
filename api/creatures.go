package api

type creatures struct {
  Name string      `json:"Name"`
  Desc description `json:"Description"`
  Img  string      `json:"Img"`
}

type description []string

func readCreatures(path string) (creatures, error) {
  var c creatures
  if err := readJSONFile(path, &c); err != nil {
    return creatures{}, err
  }
  return c, nil
}
