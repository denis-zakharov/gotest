package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDoubleHander(t *testing.T) {
	tt := []struct {
		name    string
		value   string
		doubled int
		status  int
		err     string
	}{
		{name: "double of two", value: "2", doubled: 4, status: http.StatusOK},
		{name: "missing value", value: "", status: http.StatusBadRequest, err: "missing value\n"},
		{name: "double of string", value: "string", status: http.StatusBadRequest, err: "not a number: string\n"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "localhost:8080/double?v="+tc.value, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}
			rec := httptest.NewRecorder()
			doubleHandler(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			b, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("could not read response: %v", err)
			}

			if tc.err != "" {
				if res.StatusCode != tc.status {
					t.Errorf("expected status %v; got %v", tc.status, res.StatusCode)
				}
				if msg := string(b); msg != tc.err {
					t.Errorf("expected msg %q; got %q", tc.err, msg)
				}
				return
			}
			if res.StatusCode != http.StatusOK {
				t.Fatalf("expected status 200/OK; got %v", res.StatusCode)
			}

			d, err := strconv.Atoi(string(b))
			if err != nil {
				t.Fatalf("expected an integer; got %s", b)
			}

			if d != tc.doubled {
				t.Errorf("expected double to be %v; got %v", tc.doubled, d)
			}
		})
	}
}
