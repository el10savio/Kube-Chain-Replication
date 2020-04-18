FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /app
WORKDIR /app

COPY go.mod . 
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/app


FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

ENTRYPOINT ["/go/bin/app"]

EXPOSE 8080