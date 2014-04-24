package gogl

import "testing"

const (
  ExampleUrl string = "http://www.example.com/"
  ShortUrl   string = "http://goo.gl/U98s"
)

func TestShorten(t *testing.T) {
  result, _ := Shorten(ExampleUrl)
  if result.Id != ShortUrl {
    t.Errorf("Shorten().Kind is not 'urlshortener#url': %+v", result)
  }
}

func TestExpand(t *testing.T) {
  result, _ := Expand(ShortUrl)
  if result.LongUrl != ExampleUrl {
    t.Errorf("Expand().LongUrl is not '%s': %+v", ExampleUrl, result)
  }
}
