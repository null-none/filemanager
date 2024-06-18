FROM golang:1.20.4-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main"]