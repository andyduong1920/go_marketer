.PHONY: test start-dev

build-dependencies:
	go get github.com/beego/bee

start-dev:
	bee run

test:
	go test -v -p 1 ./...
