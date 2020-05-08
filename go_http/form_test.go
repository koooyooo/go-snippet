package go_http

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestForm(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		//b, _ := ioutil.ReadAll(r.Body)
		//fmt.Println(string(b))
		e := r.ParseForm()
		if e != nil {
			panic(e)
		}
		assert.Equal(t, "World", r.Form.Get("Hello"))
		assert.Equal(t, "Bar", r.Form.Get("Foo"))
	}))
	defer ts.Close()

	vals := url.Values{}
	vals.Set("Hello", "World")
	vals.Set("Foo", "Bar")

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader(vals.Encode()))
	assert.Nil(t, err)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	defer resp.Body.Close()
}