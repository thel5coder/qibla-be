FROM qibla-base

WORKDIR /src/amqp_listener_send_notification

RUN CGO_ENABLED=0 go build -a -o /src/app


FROM alpine

RUN apk update && apk add --no-cache tzdata

WORKDIR /app

COPY --from=0 /src/app .
COPY --from=0 /src/key /key/

EXPOSE 3000

ENTRYPOINT ["./app"]
