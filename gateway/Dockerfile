FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY gateway.go ./
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
