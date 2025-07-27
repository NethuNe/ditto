# bin/bash

docker container prune -f
docker build --no-cache -t ditto -f docker/Dockerfile .  
docker run -p 1985:1985 --name ditto_main ditto