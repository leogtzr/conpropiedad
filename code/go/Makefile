.DEFAULT_GOAL := install

BIN_FILE=conpropiedad

install:
	go test
	go build -o "${BIN_FILE}"

clean:
	go clean
	rm --force conpropiedad
	rm --force cp.out

test:
	go test

check:
	go test

cover:
	go test -coverprofile cp.out
	go tool cover -html=cp.out
