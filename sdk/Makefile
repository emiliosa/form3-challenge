BINARY_NAME=sdk-account

all: build test clean

compile:
	echo "Compiling for OS [linux, freebsd] and Platform [arm64, amd64, 386]"
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/${BINARY_NAME}-freebsd-386 main.go

build:
	go build -o /app/bin/${BINARY_NAME} /app/account/account.go

test:
	go test /app/... -coverpkg=./... -v

clean:
	go clean
	rm /app/bin/${BINARY_NAME}*