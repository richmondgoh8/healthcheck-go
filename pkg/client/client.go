package client

import (
	"io/ioutil"
	"net/http"
	"time"
)

type CustomHTTP interface {
	Get(url string) ([]byte, error)
}

type CustomHTTPImpl struct {
	customClient http.Client
}

func NewClient() CustomHTTP {
	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &CustomHTTPImpl{
		customClient: *netClient,
	}
}

func (customHTTP CustomHTTPImpl) Get(url string) ([]byte, error) {
	resp, err := customHTTP.customClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
