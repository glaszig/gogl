package gogl

import "fmt"
import "log"
import "bytes"
import "net/http"
import "encoding/json"

const (
  GooglShortenUrl = "https://www.googleapis.com/urlshortener/v1/url"
  GooglExpandUrl  = "https://www.googleapis.com/urlshortener/v1/url?shortUrl=%s"
)

type ShortenRequest struct {
  LongUrl string `json:"longUrl"`
}

type GooglResponse struct {
  Kind string
  Id string
  LongUrl string
  Status string
}

func Shorten(longUrl string) GooglResponse {
  jsonStruct := &ShortenRequest{LongUrl: longUrl}
  jsonBytes, err := json.Marshal(jsonStruct)

  buf := bytes.NewBuffer(jsonBytes)
  res, err := http.Post(GooglShortenUrl, "application/json", buf)
  if err != nil {
    log.Fatal(err)
  }

  return DecodeResponse(res)
}

func Expand(shortUrl string) GooglResponse {
  expandUrl := fmt.Sprintf(GooglExpandUrl, shortUrl)

  res, err := http.Get(expandUrl)
  if err != nil {
    log.Fatal(err)
  }

  return DecodeResponse(res)
}

func DecodeResponse(res *http.Response) GooglResponse {
  var result GooglResponse
  decoder := json.NewDecoder(res.Body)
  err := decoder.Decode(&result)
  if err != nil {
    log.Fatal(err)
  }
  return result
}
