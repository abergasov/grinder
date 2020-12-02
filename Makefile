FILE_HASH := $(or ${hash},${hash},"empty_hash")
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@echo "-- building binary"
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./build/grinder ./cmd

lint:
	@echo "-- linter running"
	golangci-lint run -c .golangci.yaml ./pkg...
	golangci-lint run -c .golangci.yaml ./cmd...

test:
	@echo "-- testing pkg modules"
	go test ./pkg...

gen:
	@echo "-- generate code"
	go generate ./pkg...