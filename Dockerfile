FROM golang:1.12-alpine

ARG environment="prod"
ARG base_url="https://sessions.conference.mscc.mu"
ARG meetup_key=""
ARG meetup_secret=""
ARG jwt_secret="ThisIsNotRandomPleaseChangeIt"
ARG http_port="1337"
ARG bolt_db_path="/var/devcon.db"
ARG front_url="https://conference.mscc.mu"

ENV ENV=$environment
ENV BASE_URL=$base_url
ENV MEETUP_KEY=$meetup_key
ENV MEETUP_SECRET=$meetup_secret
ENV JWT_SECRET=$jwt_secret
ENV HTTP_PORT=$http_port
ENV BOLT_DB_PATH=$bolt_db_path
ENV FRONT_URL=$front_url

RUN apk update && apk add --virtual git && rm -rf /var/cache/apk/*

WORKDIR /opt/app

COPY . /opt/app
RUN go build -o /usr/bin/devcon-feedback

# This container exposes port 1337 to the outside world
EXPOSE $http_port

CMD ["devcon-feedback"]