.PHONY: test start-dev

start-dev:
	bee run

test:
	go test -v -p 1 ./...
