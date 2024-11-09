# bin/bash

docker build --no-cache -t ditto -f docker/Dockerfile .  
docker run -d -p 1985:1985 ditto