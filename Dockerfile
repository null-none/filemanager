FROM golang:1.20.4-alpine AS BUILDER

RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /

RUN CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server

RUN go build -o main .

CMD ["./main"]