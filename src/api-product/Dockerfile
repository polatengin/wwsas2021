FROM golang:1.16.5-buster AS build

WORKDIR /src

COPY ./ /src

RUN go build -o product .
