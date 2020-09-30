FROM golang:1.15.2
RUN go get github.com/zinderic/rpg-microservice-sandbox/quest-npc
RUN go get github.com/go-delve/delve/cmd/dlv
WORKDIR /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc
RUN go build -gcflags="all=-N -l" -o quest-npc

FROM debian:buster
WORKDIR /srv
COPY --from=0 /go/bin/dlv .
COPY --from=0 /go/src/github.com/zinderic/rpg-microservice-sandbox/quest-npc/quest-npc .
EXPOSE 2222 56268
CMD ["/srv/dlv", "--listen=:56268", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/srv/quest-npc"]
