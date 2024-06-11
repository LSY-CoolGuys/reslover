FROM golang:1.20 AS buil
WORKDIR /app
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy\
    && go build -o resolve .


FROM alpine:latest
WORKDIR /app
RUN apk update && apk add inotify-tools
COPY --from=buil /app/resolve /app/resolve
COPY --from=buil /app/monitor.sh /app/monitor.sh
RUN chmod +x /app/resolve && chmod +x /app/monitor.sh


CMD ["/bin/sh","-c","./resolve"]