#!/usr/bin/env bash

go build log-walk.go
./build_image.sh
./push_image.sh