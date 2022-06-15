package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"interview-accountapi/account"
	"interview-accountapi/internal/client"
	"interview-accountapi/mocks"
	"interview-accountapi/utils"
	"log"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("BEGIN client unit testing")

	_ = os.Setenv("ACCOUNT_API_SCHEME", "http")
	_ = os.Setenv("ACCOUNT_API_HOST", "accountapi")
	_ = os.Setenv("ACCOUNT_API_PORT", "8080")
	_ = os.Setenv("ACCOUNT_API_VERSION", "v1")

	result := m.Run()
	log.Println("END  client unit testing")
	os.Exit(result)
}

func TestNew(t *testing.T) {
	resource := "test"
	expected := MakeClient(resource)
	actual := client.New(resource)

	assert.Equal(t, expected, actual)
}

func TestClientRequest(t *testing.T) {
	t.Run("Client request GET success (with Parser mocked)", func(t *testing.T) {
		resource := "account"
		path := "test"
		method := "GET"
		config := client.GetConfig()
		parser := new(mocks.ParserMockOK)
		wrapper := client.ParserWrapper{Parser: parser}
		clientActual := client.Client{
			BaseURL: url.URL{
				Scheme: config["scheme"],
				Host:   config["host"] + ":" + config["port"],
				Path:   "/" + config["version"] + "/" + resource,
			},
			HTTPClient: http.DefaultClient,
			Parser:     wrapper,
		}

		expectedRequest, _ := http.NewRequest(method, clientActual.BaseURL.String()+"/"+path, bytes.NewBuffer(nil))
		expectedRequest.Header.Set("Accept", "application/vnd.api+json")

		actualRequest, err := clientActual.Request(method, "/"+path, nil)

		assert.NoError(t, err)
		assert.Equal(t, expectedRequest.URL, actualRequest.URL)
		assert.Equal(t, expectedRequest.RequestURI, actualRequest.RequestURI)
		assert.Equal(t, expectedRequest.Body, actualRequest.Body)
		assert.Equal(t, expectedRequest.Header, actualRequest.Header)
		assert.Equal(t, expectedRequest.Method, actualRequest.Method)
		assert.Equal(t, expectedRequest.Host, actualRequest.Host)
	})
	t.Run("Client request GET success", func(t *testing.T) {
		resource := "account"
		path := "test"
		method := "GET"
		clientActual := MakeClient(resource)

		expectedRequest, _ := http.NewRequest(method, clientActual.BaseURL.String()+"/"+path, bytes.NewBuffer(nil))
		expectedRequest.Header.Set("Accept", "application/vnd.api+json")

		actualRequest, err := clientActual.Request(method, "/"+path, nil)

		assert.NoError(t, err)
		assert.Equal(t, expectedRequest.URL, actualRequest.URL)
		assert.Equal(t, expectedRequest.RequestURI, actualRequest.RequestURI)
		assert.Equal(t, expectedRequest.Body, actualRequest.Body)
		assert.Equal(t, expectedRequest.Header, actualRequest.Header)
		assert.Equal(t, expectedRequest.Method, actualRequest.Method)
		assert.Equal(t, expectedRequest.Host, actualRequest.Host)
	})
	t.Run("Client request POST success", func(t *testing.T) {
		type Request struct {
			Data account.Account `json:"data"`
		}
		var body = Request{Data: account.MakeAccount()}
		resource := "account"
		path := "test"
		method := "POST"
		clientActual := MakeClient(resource)

		buf, _ := json.Marshal(body)
		expectedRequest, _ := http.NewRequest(method, clientActual.BaseURL.String()+"/"+path, bytes.NewBuffer(buf))
		expectedRequest.Header.Set("Accept", "application/vnd.api+json")
		expectedRequest.Header.Set("Content-Type", "application/vnd.api+json")

		actualRequest, err := clientActual.Request(method, "/"+path, body)

		assert.NoError(t, err)
		assert.Equal(t, expectedRequest.URL, actualRequest.URL)
		assert.Equal(t, expectedRequest.RequestURI, actualRequest.RequestURI)
		assert.Equal(t, expectedRequest.Body, actualRequest.Body)
		assert.Equal(t, expectedRequest.Header, actualRequest.Header)
		assert.Equal(t, expectedRequest.Method, actualRequest.Method)
		assert.Equal(t, expectedRequest.Host, actualRequest.Host)
	})
	t.Run("Client request POST fail", func(t *testing.T) {
		var body = []string{"sarasa"}
		resource := "account"
		path := "test"
		method := "POST"
		config := client.GetConfig()
		parser := new(mocks.ParserMockRequestError)
		wrapper := client.ParserWrapper{Parser: parser}
		clientActual := client.Client{
			BaseURL: url.URL{
				Scheme: config["scheme"],
				Host:   config["host"] + ":" + config["port"],
				Path:   "/" + config["version"] + "/" + resource,
			},
			HTTPClient: http.DefaultClient,
			Parser:     wrapper,
		}

		buf, err := json.Marshal(body)
		expectedRequest, err := http.NewRequest(method, clientActual.BaseURL.String()+"/"+path, bytes.NewBuffer(buf))
		expectedRequest.Header.Set("Accept", "application/vnd.api+json")
		expectedRequest.Header.Set("Content-Type", "application/vnd.api+json")

		actualRequest, err := clientActual.Request(method, "/"+path, body)

		assert.Error(t, err)
		assert.Nil(t, actualRequest)
	})
}

func TestClientResponse(t *testing.T) {
	t.Skip("TODO")
	t.Run("Cliente response suceess", func(t *testing.T) {})
}

func MakeClient(resource string) client.Client {
	config := client.GetConfig()
	parser := new(utils.Parser)
	wrapper := client.ParserWrapper{Parser: parser}

	return client.Client{
		BaseURL: url.URL{
			Scheme: config["scheme"],
			Host:   config["host"] + ":" + config["port"],
			Path:   "/" + config["version"] + "/" + resource,
		},
		HTTPClient: http.DefaultClient,
		Parser:     wrapper,
	}
}
