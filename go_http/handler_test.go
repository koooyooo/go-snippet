package go_http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerFunc(t *testing.T) {
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "host:80", r.Host)
		w.Write([]byte("Hello Body"))
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "https://host:80/path", nil)

	handler.ServeHTTP(w, r)

	assert.Equal(t, "Hello Body", w.Body.String())
}