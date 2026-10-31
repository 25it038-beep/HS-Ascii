FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /hs-ascii-live .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /hs-ascii-live .
EXPOSE 8080
CMD ["./hs-ascii-live"]
