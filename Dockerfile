##
## Build go app
##
FROM golang:1.17 AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build -o app ./cmd/api

##
## Run app
##
FROM alpine:latest

RUN apk add --no-cache \
        libc6-compat

RUN addgroup -S schattenbrot && adduser -S schattenbrot -G schattenbrot
USER schattenbrot

WORKDIR /run
COPY --from=builder /go/src/app/app ./
ENTRYPOINT [ "./app" ]