FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/mattbaird/gochimp

WORKDIR /go/src/github.com/omg-services/microservice-mandrill

ADD . /go/src/github.com/omg-services/microservice-mandrill

RUN go install github.com/omg-services/microservice-mandrill

ENTRYPOINT microservice-mandrill

EXPOSE 3000