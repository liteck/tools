package tools

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func PostForm() {

}

func PostJson(url string, jsonData []byte) (data []byte, err error) {
	var req *http.Request
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	if resp.Body == nil {
		err = errors.New("response body is null")
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	data, err = ioutil.ReadAll(resp.Body)
	return
}

func Download(url string) (data []byte, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	if resp.Body == nil {
		err = errors.New("response body is null")
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	data, err = ioutil.ReadAll(resp.Body)
	return
}
