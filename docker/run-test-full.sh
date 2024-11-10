#!/bin/bash

echo "# docker compose down"
docker compose down
echo

echo "# docker compose up -d --build test-full"
docker compose up -d --build test-full
echo

echo "# docker logs -f gower"
docker logs -f gower
echo
