package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexPage(t *testing.T) {
	assertions := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assertions.NoError(err)
	assertions.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assertions.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) {
	assertions := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	resp, err := http.Get(ts.URL)
	assertions.NoError(err)
	assertions.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine()
	assertions.NoError(err)
	assertions.Contains(string(line), "[Logger1] Started")

	line, _, err = r.ReadLine()
	assertions.NoError(err)
	assertions.Contains(string(line), "[Logger1] Completed")
}