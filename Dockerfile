FROM golang:1.13 as generator
ENV CGO_ENABLED=0
WORKDIR /src/app
COPY go.mod go.sum Makefile /src/app/
RUN make install

COPY ./ /src/app/
RUN make release

##########################

FROM alpine:latest
EXPOSE 8080

ARG DAO_COMMIT_SHA
ENV COMMIT_SHA=$DAO_COMMIT_SHA
ARG DAO_COMMIT_TAG
ENV COMMIT_TAG=$DAO_COMMIT_TAG

COPY --from=generator /src/app/dist/app /root/app

ENTRYPOINT ["/root/app"]

CMD ["serve"]
