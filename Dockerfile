FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.json config.json

CMD ["/app/main"]