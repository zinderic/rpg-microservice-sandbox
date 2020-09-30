FROM golang:1.15.2
RUN go get github.com/zinderic/rpg-microservice-sandbox/quest-npc
WORKDIR /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc
RUN go build

FROM debian:buster
WORKDIR /srv
COPY --from=0 /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc/quest-npc .
EXPOSE 2222
ENTRYPOINT ["/srv/quest-npc"]
