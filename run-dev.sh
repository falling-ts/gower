#!/bin/bash

echo "---------------- build static... ----------------"
npm run dev

echo "---------------- clean docker... ----------------"
docker compose down

echo "---------------- start dev ----------------"
docker compose up -d --build dev
