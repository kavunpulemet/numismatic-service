FROM golang:1.22.3-alpine

RUN go version

ENV GOPATH=/

COPY . .

RUN go mod download

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go build -o app ./cmd

EXPOSE 81

CMD ["./app"]