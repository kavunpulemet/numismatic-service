FROM golang:1.22.3-alpine

RUN go version

ENV GOPATH=/

COPY . .

RUN go mod download

RUN go build -o app ./cmd

EXPOSE 81

CMD ["./app"]