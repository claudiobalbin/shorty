FROM golang:1.23.1

WORKDIR /go/app

COPY . .

ENV CGO_ENABLED=0

RUN go get -d -v
RUN go build -v -o series-grpc .

EXPOSE 5000
