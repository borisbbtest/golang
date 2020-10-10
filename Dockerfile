#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY ./go/src/app/ .
#RUN apk add --no-cache git
#RUN go get -d -v ./...
RUN go install -v ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder ./go/bin/app /app
RUN ["ls","-l","./app"]
RUN ["chmod", "+x", "./app"]
ENTRYPOINT ./app
LABEL Name=golang Version=0.0.1
EXPOSE 8080
