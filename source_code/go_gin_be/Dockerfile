FROM golang:1.26.2-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o application .

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/application .

CMD ["./application"]

# docker build -t go_gin_be_img .
# docker run --name go_gin_be -d -p 8002:8080 --env-file=.env go_gin_be_img
