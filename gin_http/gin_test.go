package gin_http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGinFuncOfCtxWriter(t *testing.T) {
	// Prepare test target
	target := func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		_, err := c.Writer.WriteString("Hello Gin")
		if err != nil {
			panic(err)
		}
	}
	// Create New Context with Recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello Gin", w.Body.String())
}

func TestGinFuncOfCtxJSON(t *testing.T) {
	// Prepare test target
	target := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "Gin",
		})
	}
	// Create New Context with Recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	target(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"Hello":"Gin"}`, w.Body.String())
}

func TestServerWithGin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.Any("/path", func(c *gin.Context) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.WriteString("Hello")
	})
	// Don't call Run server this time
	// router.Run(":8080")

	// build test server with gin.Router
	ts := httptest.NewServer(router)
	defer ts.Close()

	// build real request & call
	req, err := http.NewRequest("GET", ts.URL + "/path", nil)
	assert.Nil(t, err)

	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)

	// check response from test server
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello", string(bodyBytes))
}
