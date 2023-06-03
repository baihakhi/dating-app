# Build stage
FROM golang:1.19-alpine3.16 AS build

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary ./cmd/...

# Final stage
FROM alpine:3.16

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=build /app .

# Set default port for application
EXPOSE 3030

ENTRYPOINT ["/app/binary"]
