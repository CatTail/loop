package loop

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestFSReadFile(t *testing.T) {
	loop := DefaultLoop()

	options := make(map[string]string)
	callback := func(err error, data []byte) {
		log.Printf("Execute callback with %s %s", err, data)
		assert.Nil(t, err)
		assert.Equal(t, "hello world\n", string(data))
	}
	FSReadFile(loop, "./fixtures/file.txt", options, callback)
	log.Printf("%s", loop)

	Run(loop)
}
