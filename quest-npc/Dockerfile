FROM golang:1.15.2-alpine
COPY . /go/src/quest-npc
WORKDIR /go/src/quest-npc
RUN go build

FROM alpine:latest
COPY --from=0 /go/src/quest-npc/quest-npc /srv
EXPOSE 2222
ENTRYPOINT /srv/quest-npc
