#!/bin/bash

echo "# npm run dev"
npm run dev
echo

echo "# docker compose down"
docker compose down
echo

echo "# docker compose up -d --build dev-full"
docker compose up -d --build dev-full
echo

echo "# docker logs -f gower"
docker logs -f gower
