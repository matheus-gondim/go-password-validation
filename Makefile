.PHONY: all docker-build clean

ensure-dependencies:
	go mod tidy

graphql-genereted:
	go run -mod=mod github.com/99designs/gqlgen generate

run-http:
	go run cmd/http/main.go

run-graphql:
	go run cmd/graphql/main.go

build-app: ensure-dependencies
ifeq (${OS},(Windows_NT))
	if (!(Test-Path ./build)) { mkdir build }
else
	mkdir -p build
endif
	go build -v -o ./build ./...

test:
	go clean -testcache
	go test ./...

test-cov:
ifeq (${OS},(Windows_NT))
	if (!(Test-Path ./cover)) { mkdir cover }
else
	mkdir -p cover
endif
	go test -v -coverprofile cover/cover.out ./...
	go tool cover -html cover/cover.out -o cover/cover.html

docker-build:
	VERSION_IMAGE=""
ifdef VERSION_IMAGE
	docker image build . -t go-password-validation:${VERSION_IMAGE}
else
	docker image build . -t go-password-validation
endif