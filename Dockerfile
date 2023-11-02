FROM golang:1.21.3

WORKDIR /opt/kaizen

COPY go.mod go.sum ./
RUN go mod download

COPY ./ .
RUN go build -o ./main

CMD ["./main"]
