FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/mattbaird/gochimp

WORKDIR /go/src/github.com/heaptracetechnology/microservice-mandrill

ADD . /go/src/github.com/heaptracetechnology/microservice-mandrill

RUN go install github.com/heaptracetechnology/microservice-mandrill

ENTRYPOINT microservice-mandrill

EXPOSE 3000
