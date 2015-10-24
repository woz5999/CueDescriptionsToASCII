#!/bin/bash
docker stop converter || true &&
docker rm converter || true &&
docker rmi woz5999/descriptionstoascii || true &&
./build.sh &&
docker run -d -p 80:80 --name converter woz5999/descriptionstoascii &&
sleep 2 &&
docker logs converter
