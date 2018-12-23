FROM golang:1.11

RUN mkdir /go/src/github.com

RUN mkdir /go/src/github.com/vinchauhan

RUN mkdir /go/src/github.com/vinchauhan/lockyourgate

WORKDIR /go/src/github.com/vinchauhan/lockyourgate

COPY . /go/src/github.com/vinchauhan/lockyourgate

VOLUME $PWD:/go/src/github.com/vinchauhan/lockyourgate

EXPOSE 3000

RUN go get .../.

RUN go build

CMD [ "lockyourgate" ]


