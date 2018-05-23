# Simple Dockerfile using multi-stage to limit filesize.

FROM golang:1.10
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build main.go

FROM alpine:3.7
RUN apk add --no-cache ca-certificates git
COPY --from=0 /go/src/bitbucket.code.company-name.com.au/scm/code/main .
CMD ["./main"]
