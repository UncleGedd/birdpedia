package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	assert.Nil(t, err)
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Errorf("not http ok status")
	}
	expected := "Hello World!"
	actual := recorder.Body.String()
	assert.Equal(t, expected, actual)
}

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/hello")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	respString := string(b)
	expected := "Hello World!"
	assert.Equal(t, expected, respString)
}

func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/assets/")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	contentType := resp.Header.Get("Content-Type")
	expected := "text/html; charset=utf-8"
	assert.Equal(t, expected, contentType)
}
