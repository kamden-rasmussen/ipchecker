FROM golang:1.25-alpine

WORKDIR /app

COPY . ./

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build ./cmd/main.go

CMD ["./main"]
