FROM golang:alpine AS build

# Install curl, which will be used for health checks.
RUN ["apk", "add", "curl"]

COPY ["describer.go", "/describer.go"]
RUN ["go", "build", "-o", "/describer", "/describer.go"]

ENTRYPOINT ["/describer"]
