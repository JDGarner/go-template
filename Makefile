GIT_HASH := $(shell git rev-parse --short HEAD)

dep:
	go mod download

run:
	go run main.go

lint: lint/install lint/run

lint/install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v2.5.0

lint/run:
	bin/golangci-lint run --config .golangci.yml

build:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	go build -o go-app .

docker/local-build:
	DOCKER_BUILDKIT=1 docker build -t go-app:local .

docker/ci-build:
	DOCKER_BUILDKIT=1 docker build \
	-t go-app:latest \
	-t go-app:$(GIT_HASH) .