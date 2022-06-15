package utils

import (
	"encoding/json"
	"fmt"
	errorCustom "interview-accountapi/internal/error"
	"io/ioutil"
	"net/http"
)

type Parser struct{}

func (p *Parser) ParseRequestBody(data interface{}) ([]byte, error) {
	switch data := data.(type) {
	case nil:
		return nil, nil
	case string:
		return []byte(data), nil
	default:
		buf, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		return buf, nil
	}
}

func (p *Parser) ParseResponseBody(response *http.Response, v *interface{}) error {
	responseBody, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// Status codes 200 and 201 has non-empty response body to parse
		if len(responseBody) > 0 {
			err = json.Unmarshal(responseBody, &v)
			if err != nil {
				err := fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
				return err
			}
		}

		return nil
	case http.StatusNoContent:
		// Status code 204 has empty response body
		return nil
	case http.StatusInternalServerError:
		// Status code 500 is a server error
		err := fmt.Errorf("Form3-API resource is not available: %s", response.Request.URL.String())
		return err
	default:
		// Response status not match with status 200,201,204,500
		var errorResponse errorCustom.ErrorResponse
		err := json.Unmarshal(responseBody, &errorResponse)

		if err != nil {
			return err
		}

		return errorResponse
	}
}
