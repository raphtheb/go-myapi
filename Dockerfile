# Simple Dockerfile using multi-stage to limit filesize.

FROM golang:1.10-alpine as builder

RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v  ./...
RUN go install -v ./...
RUN go build
RUN apk del git

FROM alpine:3.7
RUN apk add --no-cache ca-certificates
WORKDIR /go/src/app
COPY --from=builder /go/src/app .
CMD ["./app"]
# Be aware that running EPXOSE and -P bins the exposed container ports to ephemeral local ports.
EXPOSE 8080
