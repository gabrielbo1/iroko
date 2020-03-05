#!/bin/bash
eval "docker push gabrielbo1/iroko-$TRAVIS_CPU_ARCH;"
var amd64="amd64"
if [ "$TRAVIS_CPU_ARCH" = "$amd64" ]; then
  eval "docker push registry.heroku.com/$HEROKU_APP_NAME/web;"
  eval "heroku container:release web --app $HEROKU_APP_NAME"
fi
