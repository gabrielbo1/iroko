#!/bin/bash
eval "docker tag gabrielbo1/iroko-$TRAVIS_CPU_ARCH:${TAG} registry.heroku.com/$HEROKU_APP_NAME/web"