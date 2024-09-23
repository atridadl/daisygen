BINARY_NAME=stylegen

.PHONY: build install clean

build:
	go build -o ${BINARY_NAME}

install: build
	mv ${BINARY_NAME} ~/go/bin/

clean:
	go clean
	rm -f ${BINARY_NAME}
