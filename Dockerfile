FROM golang:1.20-alpine AS builder


WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download


COPY . .


RUN go build -o url-shortner main.go


FROM alpine:latest


WORKDIR /root/

COPY --from=builder /app/url-shortner .


EXPOSE 8000

CMD ["./url-shortner"]
