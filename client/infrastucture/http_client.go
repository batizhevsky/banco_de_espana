package infrastucture

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type HTTPClient struct {
	host   string
	prefix string
}

const (
	defaultHost   = "http://localhost:9090"
	defaultPrefix = "/"
	contentType   = "application/json"
)

var client *HTTPClient

// Post request to the bank
func (c *HTTPClient) Post(path string, json []byte) {
	resp, err := http.Post(c.url(path), contentType, bytes.NewReader(json))
	if err != nil {
		fmt.Printf("cant' create a request. Got a error: %v\n", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("cant' read a response. Got a error: %v\n", err)
	}

	fmt.Println(body)
}

func (c *HTTPClient) url(path string) string {
	return fmt.Sprintf("%s%s%s", c.host, c.prefix, path)
}

// NewHTTPClient build http client with env varialbes
func NewHTTPClient() *HTTPClient {
	if client != nil {
		return client
	}

	host, found := os.LookupEnv("BANK_HOST")
	if !found {
		host = defaultHost
	}

	prefix, found := os.LookupEnv("BANK_PREFIX")
	if !found {
		prefix = defaultPrefix
	}

	client = &HTTPClient{host, prefix}

	return client
}
