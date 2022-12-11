FROM silkeh/clang:latest

RUN apt-get update && apt-get install -y libgtest-dev libxml2 && rm -rf /var/lib/apt/lists/*

COPY . /usr/src/project
WORKDIR /usr/src/project

ENV CXX=clang++
ENV CC=clang

# CMD exec ./init.sh