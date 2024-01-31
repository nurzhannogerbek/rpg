FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd/packcalculator

RUN go build -o main .

CMD ["./main"]