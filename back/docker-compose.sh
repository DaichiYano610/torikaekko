docker-compose -f ./deployments/docker-compose.yml up -d --build
docker exec -it go_app bash
docker-compose -f ./deployments/docker-compose.yml down