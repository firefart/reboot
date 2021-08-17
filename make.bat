@echo off

SET BUILDARGS=-ldflags="-s -w" -trimpath

echo [*] Updating Dependencies
go get -u

echo [*] mod tidy
go mod tidy -v

echo [*] go fmt
go fmt ./...

echo [*] go vet
go vet ./...

echo [*] Linting
go get -u github.com/golangci/golangci-lint@master
golangci-lint run ./...
go mod tidy

rem echo [*] Tests
rem go test -v ./...

echo [*] Running build for windows
set GOOS=windows
set GOARCH=amd64
go build %BUILDARGS% -o reboot.exe
