FROM golang:1.23-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /app/main main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/main ./

RUN chmod +x /app/main

CMD ["/app/main"]