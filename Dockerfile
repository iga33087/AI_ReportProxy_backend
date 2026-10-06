# ===== Build =====
FROM golang:1.27.1-bookworm AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server .


# ===== Production =====
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=build /app/server ./server

RUN mkdir -p /app/db

EXPOSE 8080

CMD ["./server"]