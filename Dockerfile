FROM golang:1.11

RUN mkdir /go/src/app

WORKDIR /go/src/app

COPY . /go/src/app

EXPOSE 3000

RUN go get .../.

RUN go build

RUN go run main.go


