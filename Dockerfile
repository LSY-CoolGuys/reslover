FROM golang:1.20 AS buil
WORKDIR /app
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy\
    && CGO_ENABLED=0 go build -o resolve .


FROM alpine:latest
WORKDIR /app
RUN apk update && apk add inotify-tools
COPY --from=buil /app/resolve /app/resolve
COPY --from=buil /app/monitor.sh /app/monitor.sh
RUN chmod +x /app/resolve && chmod +x /app/monitor.sh \
    && mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2


CMD ["./resolve"]