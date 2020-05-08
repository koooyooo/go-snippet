package gin_http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGin(t *testing.T) {
	// Prepare test target
	target := func(c *gin.Context) {
		c.Writer.WriteString("Hello Gin")
	}

	// Create New Context with Recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target(c)

	assert.Equal(t, "Hello Gin", w.Body.String())
}
