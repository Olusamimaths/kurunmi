FROM golang:1.20.1-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN chmod +x ./start-dev.sh

CMD ["./start-dev.sh"]