FROM golang:1.17

RUN mkdir /app
WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./main.go

ENV LANG C.UTF-8
ENV CONNECTION_STRING "redis:6379"
ENV REDIS_PASSWORD ""

EXPOSE 8080

CMD ["./main"]