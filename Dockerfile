FROM golang:latest

WORKDIR /go/src/app
COPY . .

CMD ["go", "build", "LUcheck.go"]
CMD ["./LUcheck"]
