package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHander(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8080/double?v=2", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	doubleHandler(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v", res.StatusCode)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
	d, err := strconv.Atoi(string(b))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if d != 4 {
		t.Errorf("expected double to be 4; got %v", d)
	}
}
