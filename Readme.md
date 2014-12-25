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
    
    make bench

The reslut is

    BenchmarkFSReadFile       100000	    21967 ns/op
    BenchmarkFSReadFileSync	  200000	    8207 ns/op
    BenchmarkHTTPGet	      2000	        13951960 ns/op
    BenchmarkHTTPGetSync	  10	        102690509 ns/op

It seems that file read operation is so fast that extra event loop made 
asynchrounous operation slower than synchrounous one.

## License

MIT

[travis-image]: https://img.shields.io/travis/CatTail/loop.svg?style=flat
[travis-url]: https://travis-ci.org/CatTail/loop
[godoc-image]: http://img.shields.io/badge/api-Godoc-green.svg?style=flat
[godoc-url]: http://godoc.org/github.com/CatTail/loop
