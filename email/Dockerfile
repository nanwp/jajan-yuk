FROM golang:1.21-alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
COPY ../jajan-yuk.json /app/

ENV GOOGLE_APPLICATION_CREDENTIALS=/app/jajan-yuk.json

RUN go build -o email-service

FROM alpine
WORKDIR /app
COPY --from=builder /app/email-service /app
COPY --from=builder /app/jajan-yuk.json /app/

CMD ["./email-service"]
