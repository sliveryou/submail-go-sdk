package util

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Get issues a get request to the specified URL and returns the response.
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// Post issues a post form request to the specified URL and returns the response.
func Post(url string, params url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// PostMultipart issues a post multipart request to the specified URL and returns the response.
func PostMultipart(url string, params url.Values) ([]byte, error) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for key := range params {
		if key == "attachments" {
			attachments := strings.Split(params.Get(key), ",")
			if len(attachments) > 0 {
				for _, filename := range attachments {
					file, err := os.Open(filename)
					if err != nil {
						return nil, err
					}

					part, err := writer.CreateFormFile("attachments[]", filepath.Base(filename))
					if err != nil {
						return nil, err
					}

					_, err = io.Copy(part, file)
					if err != nil {
						return nil, err
					}

					_ = file.Close()
				}
			}
		} else {
			err := writer.WriteField(key, params.Get(key))
			if err != nil {
				return nil, err
			}
		}
	}

	contentType := writer.FormDataContentType()
	_ = writer.Close()

	resp, err := http.Post(url, contentType, &body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}
