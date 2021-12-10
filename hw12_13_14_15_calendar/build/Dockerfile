# Собираем в гошке
FROM golang:1.16.2 as build

ENV BIN_FILE /opt/calendar/calendar-app
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

# Кэшируем слои с модулями
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

# Собираем статический бинарник Go (без зависимостей на Си API),
# иначе он не будет работать в alpine образе.
ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/calendar/*

# На выходе тонкий образ
FROM alpine:3.9

RUN apk add --update make
COPY --from=migrate/migrate:latest /usr/bin/migrate /usr/local/bin/migrate

ENV CODE_DIR opt/calendar
WORKDIR ${CODE_DIR}

ENV BIN_FILE "/opt/calendar/calendar-app"
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE ./configs/config.toml

COPY ./.env .
COPY ./configs/config.toml ${CONFIG_FILE}
COPY ./storage ./storage
COPY ./migrations ./migrations
COPY ./Makefile .

CMD ${BIN_FILE} -config ${CONFIG_FILE}
