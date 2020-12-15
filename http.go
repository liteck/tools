package tools

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PostForm(url string, params url.Values) (data []byte, err error) {
	var resp *http.Response

	if resp, err = http.PostForm(url, params); err != nil {
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

type DownloadFile struct {
	File []byte
	Url  string
	Name string
}

func Download(url string) (file DownloadFile, err error) {
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

	file.Url = url
	file.File, err = ioutil.ReadAll(resp.Body)
	contentDisposition := resp.Header.Get("Content-Disposition")
	fN := strings.Split(contentDisposition, "filename=")
	if len(fN) == 2 && len(fN[1]) > 0 {
		file.Name = fN[1]
	}
	return
}
