FROM golang:1.22

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

EXPOSE 3001

RUN air init

CMD ["air"]