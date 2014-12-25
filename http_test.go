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
var ts = createServer()

func TestHTTPGet(t *testing.T) {
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
	loop := DefaultLoop()

	options := map[string]string{"url": ts.URL}
	callback := func(err error, data []byte) {}
	for i := 0; i < b.N; i++ {
		HTTPGet(loop, options, callback)
	}

	Run(loop)
}

func BenchmarkHTTPGetSync(b *testing.B) {
	options := map[string]string{"url": ts.URL}
	for i := 0; i < b.N; i++ {
		HTTPGetSync(options)
	}
}

func createServer() (server *httptest.Server) {
	// start test server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, content)
	}))
	return
}
