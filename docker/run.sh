#!/bin/bash

echo "# docker compose down"
docker compose down

echo "# docker compose up -d --build gower"
docker compose up -d --build gower

echo "# docker logs -f gower"
docker logs -f gower
