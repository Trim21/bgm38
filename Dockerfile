FROM alpine:latest
EXPOSE 3000

COPY ./dist/app /root/app

ENTRYPOINT ["/root/app"]
