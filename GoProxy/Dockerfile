FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get -d ./...
RUN cd $GOPATH/src/k8s.io/klog; git checkout v0.4.0
RUN go build -v main.go

CMD ["/app/main"]

EXPOSE 8090