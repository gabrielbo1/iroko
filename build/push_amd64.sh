#!/bin/bash
eval "docker push registry.heroku.com/$HEROKU_APP_NAME/web;"
eval "heroku container:release web --app $HEROKU_APP_NAME"