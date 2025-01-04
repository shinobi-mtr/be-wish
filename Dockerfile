FROM golang:1.23.4

WORKDIR /opt/
COPY . .

RUN go mod tidy
RUN go build -o out.bin . 

EXPOSE 3000

CMD ["/opt/out.bin"]
