FROM golang:1.19-alpine

COPY . /go_bootstrap

WORKDIR /go_bootstrap

ENTRYPOINT ["go", "build", "-o", "/tmp/go_bootstrap.o", "main.go"]
