@echo off

set GOOS=linux
set GOARCH=amd64
go build -o pingalive-linux

set GOOS=windows
set GOARCH=amd64
go build -o pingalive.exe

set GOOS=
set GOARCH=

echo "Build complete"
