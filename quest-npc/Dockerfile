FROM golang:1.15.2-alpine
RUN apk add git
RUN go get github.com/zinderic/rpg-microservice-sandbox/quest-npc
WORKDIR /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc
RUN go build

FROM alpine:latest
COPY --from=0 /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc/quest-npc /srv
EXPOSE 2222
ENTRYPOINT /srv/quest-npc
