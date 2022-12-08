#!/usr/bin/env bash

set -e

cd module1 && go run . && cd ..
cd module2 && go run . && cd ..
