FROM golang:1.16.3-alpine3.13
WORKDIR /go/src/api

COPY go.mod go.sum ./

RUN go get -d -v ./...
RUN go install -v ./...

COPY . .

CMD if [ "$ENV" = "development" ]; then \
        cd $(mktemp -d); \
        go mod init tmp; \
        go get -v github.com/cespare/reflex; \
        cd /go/src/api; \
        reflex -r '(\.go$|go\.mod)' -s go run cmd/datalearn/main.go; \
    else \
        go build -o main cmd/datalearn/main.go; \
        ./cmd/datalearn/main.go; \
    fi