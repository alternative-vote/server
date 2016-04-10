#!/usr/bin/env bash
docker build -t alternative-vote/server .
docker run -p 3000:3000 -v `pwd`:/usr/src/app/ -it --rm --name=server  alternative-vote/server
