#!/usr/bin/env bash

PROJECT_NAME="project"

run() {
    echo "Building an image and running a container..."
    docker build -f Dockerfile -t $PROJECT_NAME . && docker run --add-host host.docker.internal:host-gateway --net=host -it --rm -v ${PWD}:/usr/src/project $PROJECT_NAME
    # --add-host host.docker.internal:host-gateway option for macos and --net=host for linux
}

main() {
    run
    return 0
}

main $@