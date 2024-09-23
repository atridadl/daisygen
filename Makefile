BINARY_NAME=stylegen
MAIN_PATH=cmd/stylegen/main.go

.PHONY: build install clean

build:
	go build -o ${BINARY_NAME} ${MAIN_PATH}

install: build
	mv ${BINARY_NAME} ~/go/bin/

clean:
	go clean
	rm -f ${BINARY_NAME}
