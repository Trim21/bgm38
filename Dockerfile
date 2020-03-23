FROM alpine:latest
EXPOSE 3000

RUN apk add --no-cache tzdata

COPY ./dist/app /root/app

ENTRYPOINT ["/root/app"]
