FROM golang:1.13 AS builder
ENV GOPROXY=https://goproxy.cn\
    CGO_ENABLED=0
WORKDIR /usr/src/
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY / ./
RUN make

FROM alpine:latest
EXPOSE 8080
ENV GIN_MODE=release
COPY --from=builder /usr/src/dist/app /root/app
CMD ["/root/app"]