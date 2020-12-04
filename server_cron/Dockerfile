FROM qibla-base

WORKDIR /src/server_cron

RUN CGO_ENABLED=0 go build -a -o /src/app


FROM alpine

RUN apk update && apk add --no-cache tzdata

WORKDIR /app

COPY --from=0 /src/app .
COPY --from=0 /src/key /key/
COPY --from=0 /src/db /db/
COPY --from=0 /src/statics /statics/
COPY --from=0 /src/dbconfig.yml /dbconfig.yml

EXPOSE 3000

ENTRYPOINT ["./app"]
