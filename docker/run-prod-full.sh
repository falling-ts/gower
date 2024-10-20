#!/bin/bash

echo "# docker compose down"
docker compose down
echo

echo "# docker compose up -d --build prod-full"
docker compose up -d --build prod-full
echo

echo "# docker logs -f gower"
docker logs -f gower
echo
