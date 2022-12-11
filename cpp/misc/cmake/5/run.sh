#!/usr/bin/env bash

run_tests() {
    echo "Building an image and running a container..."
    docker build -f Dockerfile -t project . && docker run -it --rm -v ${PWD}:/usr/src/project project
}

main() {
    run_tests
    return 0
}

main $@