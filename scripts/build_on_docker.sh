#!/bin/bash

set -e

DOCKER_IMAGE_NAME="terraform-for-sakuraiot-build"
DOCKER_CONTAINER_NAME="terraform-for-sakuraiot-build-container"

if [[ $(docker ps -a | grep $DOCKER_CONTAINER_NAME) != "" ]]; then
  docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
fi

docker build -t $DOCKER_IMAGE_NAME .

docker run --name $DOCKER_CONTAINER_NAME \
       -e SAKURAIOT_AUTH_TOKEN \
       -e SAKURAIOT_AUTH_SECRET \
       -e TF_LOG \
       -e TESTARGS \
       $DOCKER_IMAGE_NAME make "$@"


if [[ "$@" == *"build"* ]]; then
  docker cp $DOCKER_CONTAINER_NAME:/go/src/github.com/yamamoto-febc/terraform-provider-sakuraiot/bin ./
fi
docker rm -f $DOCKER_CONTAINER_NAME 2>/dev/null
