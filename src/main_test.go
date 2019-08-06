package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)

	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok,got %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	respString := string(b)
	expected := "hello"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}
}
