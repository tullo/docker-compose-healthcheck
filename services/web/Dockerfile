FROM golang:alpine AS build

COPY ["server.go", "/server.go"]
RUN ["go", "build", "-o", "/server", "/server.go"]

ENTRYPOINT ["/server"]
