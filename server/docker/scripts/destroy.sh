containers=$(docker ps -q)
docker kill $containers
docker container rm $containers
