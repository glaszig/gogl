package gogl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const (
	GooglShortenUrl = "https://www.googleapis.com/urlshortener/v1/url"
	GooglExpandUrl  = "https://www.googleapis.com/urlshortener/v1/url?shortUrl=%s"
)

type ShortenRequest struct {
	LongUrl string `json:"longUrl"`
}

type GooglResponse struct {
	Kind    string
	Id      string
	LongUrl string
	Status  string
}

func Shorten(longUrl string) (*GooglResponse, error) {
	jsonStruct := &ShortenRequest{LongUrl: longUrl}
	jsonBytes, err := json.Marshal(jsonStruct)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(jsonBytes)
	res, err := http.Post(GooglShortenUrl, "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}

	return DecodeResponse(res)
}

func Expand(shortUrl string) (*GooglResponse, error) {
	expandUrl := fmt.Sprintf(GooglExpandUrl, url.QueryEscape(shortUrl))

	res, err := http.Get(expandUrl)
	if err != nil {
		log.Fatal(err)
	}

	return DecodeResponse(res)
}

func DecodeResponse(res *http.Response) (*GooglResponse, error) {
	result := &GooglResponse{}
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
