FROM golang:1.24-alpine AS builder

WORKDIR /usr/src/app

COPY . .
RUN go mod download
RUN go build -v

FROM alpine:3.18.4

# We'll likely need to add SSL root certificates
RUN apk --no-cache add ca-certificates

WORKDIR /usr/local/bin

COPY --from=builder /usr/src/app/app .
CMD ["./app"]
