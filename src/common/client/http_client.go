package client

import (
	"bytes"
	"errors"
	"log"
	"net/http"
)

const (
	methodGet  = "GET"
	methodPost = "POST"

	jsonContent = "application/json"
)

type Header struct {
	HeaderName  string
	HeaderValue string
}

func HTTPClient(method, url string, headers *[]Header, bodyRequest *bytes.Buffer) (resp *http.Response, err error) {

	switch method {
	case methodGet:
		resp, err = http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		return
	case methodPost:
		client := http.Client{}
		var req *http.Request
		var errClient error
		if bodyRequest != nil {
			req, errClient = http.NewRequest(method, url, bodyRequest)
		} else {
			req, errClient = http.NewRequest(method, url, nil)
		}

		if errClient != nil {
			err = errClient
			return
		}
		for _, header := range *headers {
			req.Header.Add(header.HeaderName, header.HeaderValue)
		}
		// req.Header.Add("Content-Type", "application/json")
		// req.Header.Add("nds-token", "90BH8lHyGiWJOYy54R4xmKBxDGpojCID")

		resp, err = client.Do(req)
		// if errDo != nil {
		// 	err = errDo
		// 	return
		// }
		// defer res.Body.Close()
		// resp, err = http.Post(url, jsonContent, bodyRequest)
		return
	default:
		err = errors.New("Method request is not allowed")
		return
	}

}
