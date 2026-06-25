FROM golang:1.26-alpine AS build

WORKDIR /app

COPY . ./

COPY go.mod ./
COPY go.sum ./
RUN go mod download

RUN go build ./cmd/main.go

FROM alpine:latest AS final

WORKDIR /app
COPY --from=build /app/main /app/main

CMD ["./main"]
