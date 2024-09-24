BINARY_NAME=daisygen

.PHONY: build install clean

build:
	go build -o ${BINARY_NAME} ./cmd/daisygen

install: build
	mv ${BINARY_NAME} ~/go/bin/

clean:
	go clean
	rm -f ${BINARY_NAME}
