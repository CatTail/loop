package loop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSReadFile(t *testing.T) {
	loop := DefaultLoop()

	options := make(map[string]string)
	callback := func(err error, data []byte) {
		assert.Nil(t, err)
		assert.Equal(t, "hello world\n", string(data))
	}
	FSReadFile(loop, "./fixtures/file.txt", options, callback)

	Run(loop)
}

func BenchmarkFSReadFile(b *testing.B) {
	loop := DefaultLoop()

	options := make(map[string]string)
	callback := func(err error, data []byte) {}
	for i := 0; i < b.N; i++ {
		FSReadFile(loop, "./fixtures/file.txt", options, callback)
	}

	Run(loop)
}

func BenchmarkFSReadFileSync(b *testing.B) {
	options := make(map[string]string)
	for i := 0; i < b.N; i++ {
		FSReadFileSync("./fixtures/file.txt", options)
	}
}
