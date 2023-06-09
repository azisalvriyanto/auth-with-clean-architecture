FROM golang:alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o auth-with-clean-architecture ./cmd

EXPOSE 8080

CMD ["./auth-with-clean-architecture"]
