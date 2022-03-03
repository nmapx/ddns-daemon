FROM golang:1.17-alpine

ARG HOST_UID
ARG CGO_ENABLED
ARG GOARCH
ARG GOOS

WORKDIR /go/src/github.com/nmapx/ddns-daemon
COPY . .

RUN apk --update --no-cache add -u \
        shadow \
        curl \
        make \
        less \
        bash \
        nano \
        g++ \
        git && \
    # apk del --purge \
        # git && \
    rm -rf /var/cache/apk/* \
        /tmp/*

RUN adduser -D -u ${HOST_UID} -s /bin/bash docker && \
    chown -R docker:docker .
USER docker

RUN make -f Makefile.app get

ENTRYPOINT ["./docker/app/entrypoint.sh"]
CMD ["daemon"]
