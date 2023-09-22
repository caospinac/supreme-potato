package util

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Body   interface{}
	Query  map[string]string
	Output interface{}
}

func MakeRequest(params *Request) error {
	client := new(http.Client)

	// body parsing
	var body io.Reader
	if params.Body != nil {
		bodyBytes, err := json.Marshal(params.Body)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(bodyBytes)
	}

	if params.Method == "" {
		params.Method = http.MethodGet
	}

	// request initialization
	request, err := http.NewRequest(params.Method, params.URL, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	logRequest(request)

	// query
	query := request.URL.Query()
	for queryParam, value := range params.Query {
		query.Add(queryParam, value)
	}
	request.URL.RawQuery = query.Encode()

	// request performing
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	logResponse(response, responseBody)

	if err := json.Unmarshal(responseBody, params.Output); err != nil {
		log.Panic(err)
	}

	return nil
}

func logRequest(req *http.Request) {
	log.Printf("[external_request] method:%v header:%v url:%v body:%v", req.Method, req.URL, req.Header, req.Body)
}

func logResponse(res *http.Response, responseBody []byte) {
	log.Printf("[external_response] status:%v header:%v body:%s", res.Status, res.Header, responseBody)
}
