include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build clean deploy deploy-function remove

env.dev.json:
	cp env.ex.json env.dev.json

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/telegram cmd/telegram/main.go

clean:
	rm -rf ./bin

deploy: env.dev.json clean build
	sls deploy --verbose

deploy-built: env.dev.json
	sls deploy --verbose

deploy-function: env.dev.json clean build
	sls deploy function -f ${name}

remove: clean
	sls remove --verbose
