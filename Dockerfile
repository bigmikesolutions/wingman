FROM golang:1.24-alpine as builder
ARG CMD_PATH

WORKDIR $GOPATH/src/github.com/bigmikesolutions/wingman
COPY . .

RUN go build -mod=vendor -v \
    -o /go/bin/service \
    $CMD_PATH

FROM alpine:3
RUN adduser -D wingman -u 5000
USER wingman
COPY --from=builder /go/bin/service .
CMD ["./service"]