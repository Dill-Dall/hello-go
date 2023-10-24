FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go build -o helloworld-app

EXPOSE 8080

CMD ["./helloworld-app"]

