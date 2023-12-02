FROM golang:1.21.4

WORKDIR /app/proxyservice

COPY ./proto ./proto
COPY ./proxy_service ./proxy_service
COPY ./go.* ./

RUN go work sync

WORKDIR /app/proxyservice/proxy_service
RUN go build -o proxy_service_bin .

RUN mv ./proxy_service_bin /usr/local/bin

WORKDIR /app/proxyservice

EXPOSE 50001

CMD [ "proxy_service_bin" ]