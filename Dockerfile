FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/main .
COPY --from=0 /app/migration ./migration

EXPOSE 8080

ENTRYPOINT ["./main"]