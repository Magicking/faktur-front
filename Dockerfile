FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/faktur-front
ENV PROJECT_NAME faktur-front

ADD cmd /go/src/${PROJECT_DIR}/cmd

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ENTRYPOINT /go/bin/main

EXPOSE 9090
