#!/bin/bash

WORKSPACE=$(cd $(dirname $0) && pwd -P)
#export GOROOT=/usr/local/go
export GOROOT=/usr/lib/golang
export PATH=$PATH:$GOROOT/bin
#export GOPATH=/Users/alexhan/Dev
#export GOPATH=/root/OneMarket/luck-url-go
export LUCK_URL_ENV=dev
export PROJECT_NAME=luck-url-go

if [ -n "$1" ]; then
  LUCK_URL_ENV=$1
fi
echo Environment: LUCK_URL_ENV

function build() {
  echo "Building……" && make all LUCK_URL_ENV=$LUCK_URL_ENV

  if [[ $? != 0 ]]; then
    echo -e "Build failed !"
    exit 1
  fi
  echo -e "Build success!"
}

build
