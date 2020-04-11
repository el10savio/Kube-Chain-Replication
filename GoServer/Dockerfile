FROM golang:latest

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go get -d ./...
RUN go build -v main.go

CMD ["/app/main"]

EXPOSE 8080