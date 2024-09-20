FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server

EXPOSE 3003

ENTRYPOINT ["/app/server"]

CMD ["server"]