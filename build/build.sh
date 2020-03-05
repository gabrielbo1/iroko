#!/bin/bash
eval "docker build -t gabrielbo1/iroko-$TRAVIS_CPU_ARCH:${TAG} ."
var amd64="amd64"
if [ "$TRAVIS_CPU_ARCH" = "$amd64" ]; then
  eval "docker tag gabrielbo1/iroko-$TRAVIS_CPU_ARCH registry.heroku.com/$HEROKU_APP_NAME/web"
fi