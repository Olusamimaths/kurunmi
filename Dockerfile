FROM golang:1.20.1-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

COPY wait-for-it.sh /usr/wait-for-it.sh
RUN chmod +x /usr/wait-for-it.sh

RUN chmod +x ./start-dev.sh

CMD ["./start-dev.sh"]