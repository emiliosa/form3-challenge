package mocks

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
)

type ParserMockOK struct {
	mock.Mock
}

type ParserMockRequestError struct {
	mock.Mock
}

func (w *ParserMockOK) ParseRequestBody(data interface{}) ([]byte, error) {
	return nil, nil
}

func (w *ParserMockOK) ParseResponseBody(response *http.Response, v *interface{}) error {
	return nil
}

func (w *ParserMockRequestError) ParseRequestBody(data interface{}) ([]byte, error) {
	return nil, errors.New("this is an error")
}

func (w *ParserMockRequestError) ParseResponseBody(response *http.Response, v *interface{}) error {
	return errors.New("this is an error")
}
