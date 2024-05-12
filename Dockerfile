##
## Build
##
FROM golang:1.22 as build
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./

ADD . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o build/ ./cmd/...

##
## Deploy
##
FROM alpine:3
ENV APP_HOME /app
ENV USER "app"
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
RUN addgroup -g 1000 -S "${USER}" && adduser -s /bin/sh -u 1000 -G "${USER}" -D "${USER}"

USER "${USER}"

COPY --from=build --chown="${USER}":"${USER}" /app/build "$APP_HOME"
CMD ["/app/cmd"]
