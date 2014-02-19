package main

import (
  "bytes"
  "errors"
  "net/http"
)

func Fetch(url string) ([]byte, error) {
  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    return nil, errors.New(resp.Status)
  }
  buffer := new(bytes.Buffer)
  buffer.ReadFrom(resp.Body)
  return buffer.Bytes(), nil
}
