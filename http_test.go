package loop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var content = "Hello, World"

func TestHTTPGet(t *testing.T) {
	ts := createServer()
	defer ts.Close()

	loop := DefaultLoop()

	options := map[string]string{"url": ts.URL}
	callback := func(err error, data []byte) {
		assert.Nil(t, err)
		assert.Equal(t, content, string(data))
	}
	HTTPGet(loop, options, callback)

	Run(loop)
}

func BenchmarkHTTPGet(b *testing.B) {
	ts := createServer()
	defer ts.Close()

	loop := DefaultLoop()

	options := map[string]string{"url": ts.URL}
	callback := func(err error, data []byte) {}
	for i := 0; i < b.N; i++ {
		HTTPGet(loop, options, callback)
	}

	Run(loop)
}

func BenchmarkHTTPGetSync(b *testing.B) {
	ts := createServer()
	defer ts.Close()

	options := map[string]string{"url": ts.URL}
	for i := 0; i < b.N; i++ {
		HTTPGetSync(options)
	}
}

func createServer() (server *httptest.Server) {
	// start test server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1000)
		fmt.Fprint(w, content)
	}))
	return
}
