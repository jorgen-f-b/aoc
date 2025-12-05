#!/bin/bash

pushd $1 > /dev/null

gcc main.c -o main
./main
rm main

popd > /dev/null
