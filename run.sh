#!/bin/bash

echo "---------------- 清理容器 ----------------"
docker compose down

echo "---------------- 启动测试 ----------------"
docker compose up -d --build gower
