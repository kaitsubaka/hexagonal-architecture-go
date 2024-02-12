FROM golang:1.21.6-alpine as builder

ENV APP_ENV=develpment
ENV APP_PORT=3003

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod verify

RUN go build -o /app/server ./cmd/main.go

EXPOSE $APP_PORT

CMD ["/app/server"]