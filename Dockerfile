# Stage 1: Build the Go binary
FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/url-shortner

# Stage 2: Create a minimal runtime image
FROM alpine

RUN adduser -S -D -H -h /app appuser

USER appuser

WORKDIR /app

COPY --from=builder /build/main /app/

EXPOSE 3000

CMD ["./main"]
