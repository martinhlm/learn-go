package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*func TestHandler(t *testing.T) {

	// table driven test
	cases := []struct {
		in, out string
	}{
		{"m@golang.org", "gopher m"},
		{"something", "dear something"},
	}

	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)

		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status 200, got: %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}
}*/

func BenchmarkHandler(b *testing.B) {

	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/m@golang.org",
			nil,
		)

		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}

		rec := httptest.NewRecorder()
		handler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected status 200, got: %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher m") {
			b.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}
}
