FROM golang:alpine AS builD
WORKDIR /app
ADD . /app
RUN cd /app \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy\
    && go build -o resolve .


FROM alpine:latest
WORKDIR /app
RUN apk update && apk add inotify-tools
COPY --from=builD /app/resolve /app/resolve
COPY --from=builD /app/monitor.sh /app/monitor.sh
RUN chmod +x /app/resolve
RUN chmod +x /app/monitor.sh
RUN cd /app

CMD ["/bin/sh","-c","./monitor.sh && ./resolve"]