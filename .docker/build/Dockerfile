FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp