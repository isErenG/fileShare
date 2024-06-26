#!/bin/bash

echo "Shutting down containers"
docker-compose down

# Prune unused containers
echo "Pruning unused containers..."
docker container prune -f

# Prune unused images
echo "Pruning unused images..."
docker image prune -a -f

# Prune unused volumes
echo "Pruning unused volumes..."
docker volume prune -f

# Prune unused networks
echo "Pruning unused networks..."
docker network prune -f

# Confirm completion
echo "Docker system prune complete!"
