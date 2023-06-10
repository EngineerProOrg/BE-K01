FROM golang:1.14

ENV GO111MODULE=on

WORKDIR /app

RUN mkdir /var/log/entry/
RUN go mod download
COPY ../../ .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build cmd/authen_and_post_svc/main.go

CMD ["/app/main -config config/files/test.yml"]