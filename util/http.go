package util

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestOptions struct {
	Url     string
	Method  string
	Body    io.Reader
	Query   map[string]string
	Headers map[string]string
}

func Request(options *RequestOptions, result interface{}) (
	err error,
) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Method, options.Url, options.Body)
	if err != nil {
		return
	}

	if options.Query != nil {
		q := req.URL.Query()
		for index, value := range options.Query {
			q.Add(index, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	if options.Headers != nil {
		for prop, value := range options.Headers {
			req.Header.Add(prop, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	Body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(Body, result)
	if err != nil {
		return
	}
	return
}

func RawRequest(options *RequestOptions) (
	*http.Response,
	error,
) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Method, options.Url, options.Body)
	if err != nil {
		return nil, err
	}

	if options.Query != nil {
		q := req.URL.Query()
		for index, value := range options.Query {
			q.Add(index, value)
		}

		req.URL.RawQuery = q.Encode()
	}

	if options.Headers != nil {
		for prop, value := range options.Headers {
			req.Header.Add(prop, value)
		}
	}

	return client.Do(req)
}
