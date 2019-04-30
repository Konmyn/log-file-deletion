#!/usr/bin/env bash

export $(cat .env | xargs)

docker build -t "$image" .
