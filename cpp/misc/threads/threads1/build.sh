#!/usr/bin/env bash

if [ ! -x build ]; then
    mkdir build
fi
cd build
rm -fr *
cmake ..
make -sj4
echo 
echo
echo
./project
cd ..