package client

import (
	"bytes"
	"interview-accountapi/utils"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseURL    url.URL
	HTTPClient *http.Client
	Parser     ParserWrapper
}

type Parser interface {
	ParseRequestBody(data interface{}) ([]byte, error)
	ParseResponseBody(response *http.Response, v *interface{}) error
}

type ParserWrapper struct {
	Parser Parser
}

func GetConfig() map[string]string {
	config := make(map[string]string)
	config["scheme"] = os.Getenv("ACCOUNT_API_SCHEME")
	config["host"] = os.Getenv("ACCOUNT_API_HOST")
	config["port"] = os.Getenv("ACCOUNT_API_PORT")
	config["version"] = os.Getenv("ACCOUNT_API_VERSION")

	return config
}

func New(resource string) Client {
	config := GetConfig()
	parser := new(utils.Parser)
	wrapper := ParserWrapper{Parser: parser}
	return Client{
		BaseURL: url.URL{
			Scheme: config["scheme"],
			Host:   config["host"] + ":" + config["port"],
			Path:   "/" + config["version"] + "/" + resource,
		},
		HTTPClient: http.DefaultClient,
		Parser:     wrapper,
	}
}

func (client *Client) Request(method string, path string, data interface{}) (*http.Request, error) {
	body, err := client.Parser.Parser.ParseRequestBody(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, client.BaseURL.String()+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if data != nil {
		request.Header.Set("Content-Type", "application/vnd.api+json")
	}

	request.Header.Set("Accept", "application/vnd.api+json")
	return request, nil
}

func (client *Client) Do(request *http.Request, v interface{}) (*http.Response, error) {
	response, err := client.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response, client.Parser.Parser.ParseResponseBody(response, &v)
}
