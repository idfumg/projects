#!/usr/bin/env bash

PROJECT_NAME="project"

run() {
    echo "Building an image and running a container..."
    docker build -f Dockerfile -t $PROJECT_NAME . && docker run -it --rm -v ${PWD}:/usr/src/project $PROJECT_NAME
}

main() {
    run
    return 0
}

main $@