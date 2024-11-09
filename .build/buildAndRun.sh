# bin/bash

go build -o ditto
docker build -t ditto ./docker
docker run -d -p 1985:1985 ditto