FROM golang:1.15.2-alpine3.12

ENV GIN_MODE=release
 
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor