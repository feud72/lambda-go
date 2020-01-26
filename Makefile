.PHONY: build clean deploy

build:
	env DEPNOLOCK=1 dep ensure -v
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/handler.go
	chmod 777 bin/handler
	chmod 777 bin

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
