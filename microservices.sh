#!/usr/bin/env bash

microservices=( quest-npc world story-teller )

start() {
    echo 'Starting all microservices..'
    for service in ${microservices[*]}
    do
        ( cd $service && skaffold run )
    done
}

stop() {
    echo 'Stopping all microservices..'
    for service in ${microservices[*]}
    do
        ( cd $service && skaffold delete )
    done
}

case "$1" in
    start) start ;;
    stop) stop ;;
    *) echo 'Please use correct subcommand - start or stop' ;;
esac

