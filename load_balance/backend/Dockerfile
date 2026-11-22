FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]

# docker build -t img-backend .
# docker run --hostname backend1 --name backend -p 8080:8080 -d img-backend
# docker run --hostname backend2 --name backend -p 8080:8080 -d img-backend
# docker run --hostname backend3 --name backend -p 8080:8080 -d img-backend