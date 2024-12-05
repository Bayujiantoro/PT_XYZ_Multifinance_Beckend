FROM golang:1.21-alpine

WORKDIR / app

COPY go.mod go.sum ./

