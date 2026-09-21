FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY go_service.go ./
RUN go build -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8002
CMD ["./main"]
