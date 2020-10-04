FROM golang:1.15.2
RUN go get github.com/zinderic/rpg-microservice-sandbox/world
WORKDIR /go/src/github.com/zinderic/rpg-microservice-sandbox/world
RUN go build

FROM debian:buster
WORKDIR /srv
COPY --from=0 /go/src/github.com/zinderic/rpg-microservice-sandbox/world/world .
EXPOSE 2222
ENTRYPOINT ["/srv/world"]
