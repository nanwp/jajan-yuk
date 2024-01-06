# Dockerfile di service1
FROM golang:1.21-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

RUN go build -o email-service

FROM alpine
WORKDIR /app
COPY --from=builder /app/email-service /app

CMD ["./email-service"]
