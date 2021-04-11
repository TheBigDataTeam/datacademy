FROM golang:1.16.3-alpine3.13
WORKDIR /api
COPY . .
RUN go build -o main main.go

EXPOSE 3100

CMD ["/api/main"]