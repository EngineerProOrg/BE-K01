FROM golang:1.19.8 AS builder

ENV GO111MODULE=on
WORKDIR /app

RUN mkdir /var/log/entry/
COPY ../../ .
RUN --mount=type=cache,target=/go/pkg/mod/cache go mod download

RUN --mount=type=cache,target=/go/pkg/mod/cache CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o newsfeed_svc.linux -mod=mod cmd/newsfeed_svc/main.go

CMD ["/app/newsfeed_svc.linux"]

FROM builder AS test_env

RUN cp /app/configs/files/test.yml /app/config.yml
CMD ["/app/newsfeed_svc.linux"]

FROM builder AS live_env
RUN cp /app/configs/files/live.yml /app/config.yml
CMD ["/app/newsfeed_svc.linux"]