package go_http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestingHandlerFunc(t *testing.T) {
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// check request
		assert.Equal(t, "host:80", r.Host)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello Body"))
	})
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "https://host:80/path", nil)

	handler.ServeHTTP(w, r)

	// assert response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Body", w.Body.String())
}