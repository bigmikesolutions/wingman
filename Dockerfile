FROM golang:1.17-alpine as builder
ARG CMD_PATH
ARG VER
WORKDIR $GOPATH/src/github.com/bigmikesolutions/wingman
COPY . .
RUN go build -mod=vendor -v \
    -ldflags "-X github.com/bigmikesolutions/wingman/pkg/build.Version=${VER}" \
    -o /go/bin/service \
    $CMD_PATH

FROM alpine:3
RUN adduser -D wingman -u 5000
USER wingman
COPY --from=builder /go/bin/service .
CMD ["./service"]