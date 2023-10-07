
FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go
RUN go build -o consumer consumer.go

EXPOSE 6000

CMD ["./main", "./consumer"]