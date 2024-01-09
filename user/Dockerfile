it FROM golang:1.21-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .

RUN go build -o user-service

FROM alpine
WORKDIR /app
COPY --from=builder /app/user-service /app

CMD ["./user-service"]
