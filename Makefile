.DEFAULT_GOAL := build

BIN_FILE=owenmoney

build:
	@go build -o "${BIN_FILE}"
	@ls -l "${BIN_FILE}"

clean:
	go clean
	rm --force "${BIN_FILE}"

test:
	go test

run:
	./"${BIN_FILE}"
