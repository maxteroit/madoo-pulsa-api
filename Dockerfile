FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o madoo-pulsa-api ./cmd/main.go

EXPOSE 8082
CMD ["./madoo-pulsa-api"]