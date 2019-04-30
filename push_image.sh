#!/usr/bin/env bash

export $(cat .env | xargs)

docker push "$image"
