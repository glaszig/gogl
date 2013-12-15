package gogl

import "testing"

func TestMain(t *testing.T) {
  result := shorten("http://www.example.com")
  if result.Kind != "urlshortener#url" {
    t.Errorf("shorten().Kind is not 'urlshortener#url': %+v", result)
  }
}
