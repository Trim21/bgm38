FROM alpine:latest
RUN apk add --no-cache tzdata

COPY ./dist/app /root/app

ENTRYPOINT ["/root/app"]
