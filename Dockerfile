FROM golang:1.20

WORKDIR $GOPATH/src/github.com/jadilet/calculator

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o ./out/calculator .

ENTRYPOINT ["./out/calculator"]