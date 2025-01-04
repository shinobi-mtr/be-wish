FROM golang:1.23.4-alpine AS builder

WORKDIR /opt/
COPY . .

RUN go mod tidy
RUN go build -ldflags="-extldflags=-static" -o out.bin . 

FROM alpine:3.18.2

COPY --from=builder /opt/out.bin /opt/out.bin

EXPOSE 3000

CMD ["/opt/out.bin"]
