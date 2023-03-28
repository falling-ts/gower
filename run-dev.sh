#!/bin/bash

echo "---------------- build static... ----------------"
npm run dev

echo "---------------- start dev ----------------"
docker compose up -d dev
