FROM alpine:latest
EXPOSE 8080
COPY ./dist/app /root/app

ENTRYPOINT ["/root/app"]

CMD ["serve"]
