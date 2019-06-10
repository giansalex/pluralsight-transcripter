# Contributing

## Requirements
- Install MinGW (`choco install mingw`).
- Install [dep](https://golang.github.io/dep/).

## Build
Clone repository, install dependencies and build.

```
git clone https://github.com/giansalex/pluralsight-transcripter.git
cd pluralsight-transcripter
dep ensure
go build
```

## Release

For x64
```
GOOS=windows GOARCH=amd64 windres -o main-res.syso main.rc && go build -ldflags "-s -w" -a -installsuffix cgo -o main.exe
```