package client

import (
	"io"
	"net/http"
)

type ResponseDTO struct {
	Status      string
	StatusCode  int
	ContentType string
	Body        []byte
}

func Send(request *http.Request) (resp ResponseDTO, err error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()

	status := response.Status
	statusCode := response.StatusCode
	contentType := response.Header.Get("Content-Type")

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return resp, err
	}

	resp = ResponseDTO{Status: status, StatusCode: statusCode, ContentType: contentType, Body: body}

	return resp, err
}
