package gogl

import "log"
import "fmt"
import "bytes"
import "net/http"
import "encoding/json"

const GOOGLE_BASE_URL string = "https://www.googleapis.com/urlshortener/v1"

type GooglResponse struct {
  Kind string
  Id string
  LongUrl string
}

func shorten(longUrl string) GooglResponse {
  json := fmt.Sprintf("{\"longUrl\":\"%s\"}", longUrl)
  url  := fmt.Sprintf("%s/%s", GOOGLE_BASE_URL, "url")

  buf := bytes.NewBufferString(json)
  res, err := http.Post(url, "application/json", buf)
  if err != nil {
    log.Fatal(err)
  }

  return decode(res)
}

func client() http.Client {
  return http.Client{}
}

func decode(res *http.Response) GooglResponse {
  var result GooglResponse
  decoder := json.NewDecoder(res.Body)
  err := decoder.Decode(&result)
  if err != nil {
    log.Fatal(err)
  }
  return result
}
