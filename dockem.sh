#!/bin/bash

set -e

USER=cmar82
IMAGE=go-chi-api
VERSION=1.0.2

main () {
    docker build -t $USER/$IMAGE:$VERSION .
    docker push $USER/$IMAGE:$VERSION
}

main "$@"