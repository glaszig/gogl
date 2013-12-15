package gogl

import "log"
import "bytes"
import "net/http"
import "encoding/json"

const GOOGL_SHORTEN_URL string = "https://www.googleapis.com/urlshortener/v1/url"

type ShortenRequest struct {
  LongUrl string `json:"longUrl"`
}

type GooglResponse struct {
  Kind string
  Id string
  LongUrl string
}

func shorten(url string) GooglResponse {
  jsonStruct := &ShortenRequest{LongUrl: url}
  jsonBytes, err := json.Marshal(jsonStruct)

  log.Printf("Shortening request: %s\n", string(jsonBytes))

  buf := bytes.NewBuffer(jsonBytes)
  res, err := http.Post(GOOGL_SHORTEN_URL, "application/json", buf)
  if err != nil {
    log.Fatal(err)
  }

  return decode(res)
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
