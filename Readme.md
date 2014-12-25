# loop

[![Build Status][travis-image]][travis-url]
[![API Documentation][godoc-image]][godoc-url]

> Event loop for golang. Just for fun!

## Installation

    go get github.com/CatTail/loop

## Usage

	loop := DefaultLoop()

	options := map[string]string{"url": "http://google.com"}
	callback := func(err error, data []byte) {
		assert.Nil(t, err)
		assert.Equal(t, content, string(data))
	}
	HTTPGet(loop, options, callback)

	Run(loop)

## Benchmark

Run
    
    go test -bench=.

To execute golang benchmark

    BenchmarkFSReadFile	        300	   4266092 ns/op
    BenchmarkFSReadFileSync	    500	   3364286 ns/op
    BenchmarkHTTPGet	        2000   15359451 ns/op
    BenchmarkHTTPGetSync	    10	   103618460 ns/op

It seems that file read operation is so fast that extra event loop made 
asynchrounous operation slower than synchrounous one.

Run

    node ./benchmark.js

To execute Node.js benchmark

    Benchmark readFileSync  300     57137028.3762614ns/op
    Benchmark readFile      300     21750434.239705402ns/op

In Node.js, the benchmark result is different, asynchrounous operation is
faster than the synchrounous one.

## License

MIT

[travis-image]: https://img.shields.io/travis/CatTail/loop.svg?style=flat
[travis-url]: https://travis-ci.org/CatTail/loop
[godoc-image]: http://img.shields.io/badge/api-Godoc-green.svg?style=flat
[godoc-url]: http://godoc.org/github.com/CatTail/loop
