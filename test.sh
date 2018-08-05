#!/bin/bash
docker rm -f converter || true &&
make &&
docker run -d -p 80:80 --name converter woz5999/cuedescriptionstoascii &&
sleep 2 &&
docker logs converter
