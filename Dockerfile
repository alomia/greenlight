FROM golang:1.21.3

WORKDIR /usr/src/app

COPY /bin/linux_amd64/greenlight ./

COPY migrations ./migrations

EXPOSE 4000

ENV GREENLIGHT_DB_DSN=''

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

RUN mv migrate.linux-amd64 $GOPATH/bin/migrate

# RUN migrate -path ./migrations -database $GREENLIGHT_DB_DSN up

CMD [ "./greenlight" ]
