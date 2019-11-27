FROM golang

RUN go get github.com/gorilla/mux

RUN go get github.com/mattbaird/gochimp

WORKDIR /go/src/github.com/oms-services/mandrill

ADD . /go/src/github.com/oms-services/mandrill

RUN go install github.com/oms-services/mandrill

ENTRYPOINT mandrill

EXPOSE 3000
