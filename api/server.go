package api

type config struct {
  Port uint `json:"port"`
}

func readConfig(path string) (config, error) {
  var c config
  if err := readJSONFile(path, &c); err != nil {
    return config{}, err
  }
  return c, nil
}
