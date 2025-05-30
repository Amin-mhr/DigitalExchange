FROM golang:alpine AS builder
LABEL maintainer="Amin Moharrami"
# ENVs
ENV CGO_ENABLED 0 \
    GOPATH /go \
    GOCACHE /go-build \
    GOOS linux \
    GOARCH=amd64 \
    GO111MODULE=on
# Copy
WORKDIR /digitalExchange
COPY . /digitalExchange
# Install Dependencies
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go mod download

# Build the binary
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go build -o digitalExchange main.go

CMD ["./digitalExchange", "app", "bootstrap"]
