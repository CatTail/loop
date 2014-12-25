test:
	go test -v ./...

bench:
	go test -bench=.
	node ./benchmark.js

clean:
	rm -rf build/*

build: clean
	gox -os="linux darwin" -output="build/{{.Dir}}_{{.OS}}_{{.Arch}}"
	find build -type f | xargs -I_file -- sh -c 'tar czvf _file.tar.gz _file && rm _file'

get:
	go get github.com/tools/godep
	godep restore ./...

.PHONY: test bench clean build get
