FROM golang:1.21.4

WORKDIR /app/ticketservice

COPY ./proto ./proto
COPY ./ticket_service ./ticket_service
COPY ./go.* ./

RUN go work sync

WORKDIR /app/ticketservice/ticket_service
RUN go build -o ticket_service_bin .

RUN mv ./ticket_service_bin /usr/local/bin

WORKDIR /app/ticketservice

RUN ls

EXPOSE 50051

CMD [ "ticket_service_bin" ]