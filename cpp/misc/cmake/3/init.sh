#!/usr/bin/env bash

mkdir build
cd build
cmake ..
make -sj4 test_regression
ls
./test_regression