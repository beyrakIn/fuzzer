build:
	GOOS=windows GOARCH=amd64 go build -o bin/fuzzer-windows64.exe
	GOOS=linux GOARCH=amd64 go build -o bin/fuzzer-linux64
	GOOS=darwin GOARCH=amd64 go build -o bin/fuzzer-macos64
