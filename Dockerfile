FROM golang:1.11-alpine
ADD . /go/src/github.com/Percona-Lab/pxc-service-broker
EXPOSE 8081
WORKDIR /Users/nonemax/go/src/github.com/Percona-Lab/pxc-service-broker/cmd/broker
RUN go install

ENTRYPOINT ["broker"]