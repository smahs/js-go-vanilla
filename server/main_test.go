package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	log.Println("Starting the mock server...")
	Version = "1.1.1"
}

func shutdown() {
	log.Println("Shutting the mock server...")
	// Nothing to purge, continue exiting...
}

func testCommon(url string, h http.HandlerFunc) (*http.Response, error) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", url, nil)
	h(w, r)

	if w.Code != 200 {
		return nil, fmt.Errorf(
			"Expected code 200; received: %d", w.Code,
		)
	}

	resp := w.Result()
	return resp, nil
}

func TestFormatVersion(t *testing.T) {
	if !strings.HasSuffix(formatVersion(), Version) {
		t.Errorf("Version mismatch")
	}
}

func TestVersion(t *testing.T) {
	var resp, err = testCommon("/version", serveVersion)
	if err != nil {
		t.Errorf(err.Error())
	}

	if resp.Body == nil {
		t.Errorf("Expected a non-empty response body")
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if !strings.HasSuffix(string(body), Version) {
		t.Errorf("Version mismatch")
	}
}

func TestStatic(t *testing.T) {
	var resp, err = testCommon("/index.html", serveStatic)
	if err != nil {
		t.Errorf(err.Error())
	}

	log.Println(resp.Header)
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		t.Errorf("Expected \"text/html\" content-type header")
	}
}
