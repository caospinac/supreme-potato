package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
	Offset  *int64      `json:"offset,omitempty"`
	Limit   *int64      `json:"limit,omitempty"`
}

type ResponseBuilder interface {
	WithBody(interface{}) ResponseBuilder
	WithHeader(string, string) ResponseBuilder
	WithStatus(int) ResponseBuilder
	WithMessage(string) ResponseBuilder

	build() EventResponse
}

type responseBuilder struct {
	eventResponse EventResponse

	status  int
	message string
	body    interface{}
	headers map[string]string
}

func NewResponse() ResponseBuilder {
	return &responseBuilder{
		status: http.StatusOK,
		headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func (b *responseBuilder) WithBody(in interface{}) ResponseBuilder {
	b.body = in

	return b
}

func (b *responseBuilder) WithHeader(name, value string) ResponseBuilder {
	b.eventResponse.Headers[name] = value

	return b
}

func (b *responseBuilder) WithStatus(code int) ResponseBuilder {
	b.status = code

	return b
}

func (b *responseBuilder) WithMessage(message string) ResponseBuilder {
	b.message = message

	return b
}

func (b *responseBuilder) build() EventResponse {
	if b.message == "" {
		b.message = http.StatusText(b.status)
	}

	responseBody := &Response{
		Status:  b.status,
		Message: b.message,
	}

	if b.body != nil {
		responseBody.Body = b.body
	}

	responseBytes, err := json.Marshal(responseBody)
	if err != nil {
		panic(err)
	}

	var responseBuffer bytes.Buffer
	json.HTMLEscape(&responseBuffer, responseBytes)

	return &events.APIGatewayV2HTTPResponse{
		StatusCode:      b.status,
		IsBase64Encoded: false,
		Headers:         b.headers,
		Body:            responseBuffer.String(),
	}
}

func whoami(skip int) string {
	pc, _, _, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown"
	}

	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}

	return me.Name()
}
