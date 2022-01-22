FROM golang:1.17 as builder

WORKDIR /root/test-service
COPY . .
RUN go build cmd/server/main.go

FROM ubuntu:18.04

WORKDIR /root
COPY --from=builder /root/test-service/main ./
CMD ["./main"]
