FROM golang:1.17 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLE = 0 go build -o vending cmd

FROM alpine:3.15.3

WORKDIR /app

COPY --from=builder  /app/cmd/vending .

CMD ["/vending"]