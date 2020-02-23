#!/bin/bash
docker push registry.heroku.com/$HEROKU_APP_NAME/web;
heroku container:release web --app $HEROKU_APP_NAME