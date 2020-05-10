package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

func GetCurlCommand(req http.Request) (string, error) {
	xcommand := []string{"curl"}
	xcommand = append(xcommand, "-X")
	xcommand = append(xcommand, bashEscape(req.Method))

	if req.Body != nil {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return "", err
		}
		req.Body = nopCloser{bytes.NewBuffer(body)}
		bodyEscaped := bashEscape(string(body))
		xcommand = append(xcommand, "-d")
		xcommand = append(xcommand, bodyEscaped)
	}

	var keys []string

	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		xcommand = append(xcommand, "-H")
		xcommand = append(xcommand, bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}

	xcommand = append(xcommand, bashEscape(req.URL.String()))

	return strings.Join(xcommand, " "), nil
}

// nopCloser is used to create a new io.ReadCloser for req.Body
type nopCloser struct {
	io.Reader
}

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

func (nopCloser) Close() error { return nil }
