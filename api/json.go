package api

import (
  "encoding/json"
  "io/ioutil"
)

func readJSONFile(path string, output interface{}) error {
  data, err := ioutil.ReadFile(path)
  if err != nil {
    return err
  }
  if err := json.Unmarshal(data, output); err != nil {
    return err
  }
  return nil
}
