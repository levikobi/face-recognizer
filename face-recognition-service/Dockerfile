FROM golang:1.16-alpine AS build_base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -o ./out/go-sample-app -mod=vendor .

# Start fresh from a smaller image
FROM alpine:3.14
RUN apk add ca-certificates

COPY --from=build_base /tmp/go-sample-app/out/go-sample-app /app/go-sample-app

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
RUN chmod +x /app/go-sample-app
CMD ["/app/go-sample-app"]

