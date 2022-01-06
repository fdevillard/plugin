package plugin

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	hijackHttptest "github.com/getlantern/httptest"
)

// This test ensures that the middleware supports Hijacked response handler. More precisely, a response handler may be
// Hijacked by the other handlers (see: https://pkg.go.dev/net/http#Hijacker). This typically occurs for websocket connection.
// The purpose of hijacked response handler is to have fine grain control of the answer returned to the user.
func TestWrappingHijackedHandler(t *testing.T) {
	const expectContent = "Been there, got hijacked."

	// Prepare hijacking handler
	hijackHandler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		hijack, ok := rw.(http.Hijacker)
		if !ok {
			t.Errorf("failed to hijack the writer: %+v", rw)
		}
		connection, _, err := hijack.Hijack()
		if err != nil {
			t.Errorf("couldn't hijack the connection due to error: %v", err)
		}
		defer connection.Close()

		if _, err := connection.Write([]byte(expectContent)); err != nil {
			t.Errorf("fail to write to connection due to error: %v", err)
		}
	})

	// Prepare the audit handler
	config := &Config{}
	handler, _ := New(context.Background(), hijackHandler, config, "test")

	// Prepare an request with both types of headers
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("shouldn't be used"))

	// Run the query
	w := hijackHttptest.NewRecorder(nil)
	handler.ServeHTTP(w, req)

	// Assert response
	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("failed to read content due to error: %v", err)
	}
	bodyStr := string(body)
	if bodyStr != expectContent {
		t.Errorf("received content (%s) differ from the expected one (%s)", bodyStr, expectContent)
	}
}
