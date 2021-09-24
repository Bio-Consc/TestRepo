FROM golang:1.13

WORKDIR /src

COPY . .

RUN GO111MODULE=on go build -o application

ENTRYPOINT ./application