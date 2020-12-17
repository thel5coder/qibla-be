FROM qibla-base

WORKDIR /src/server

RUN CGO_ENABLED=0 go build -a -o /src/app


FROM alpine

RUN apk update && apk add --no-cache tzdata postgresql-client wkhtmltopdf ttf-ubuntu-font-family

WORKDIR /app

COPY --from=0 /src/app .
COPY --from=0 /src/key /key/
COPY --from=0 /src/db /db/
COPY --from=0 /src/statics /statics/
COPY --from=0 /src/dbconfig.yml /dbconfig.yml

EXPOSE 3000

ENTRYPOINT ["./app"]
