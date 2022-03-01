#!/bin/bash

set -e

USER=ensooboka
IMAGE=go-chi-api
VERSION=1.0.0

main () {
    docker build -t $USER/$IMAGE:$VERSION .
    docker push $USER/$IMAGE:$VERSION
}

main "$@"