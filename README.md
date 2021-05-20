# gosh

This CLI tool is used to ssh to multiple machines at once

# Quickstart Instructions

Install the Library
```go
go get github.com/dougkirkley/gosh
```

Install the binary for your OS

in the pkg dir ex: ~/go/src/github.com/dougkirkley/gosh/
```go
go build
```

Example

```sh
./gosh -H 10.0.0.1 -c "whoami"
```