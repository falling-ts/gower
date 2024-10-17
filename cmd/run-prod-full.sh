#!/bin/bash

echo "---------------- 清理容器 ----------------"
docker-compose down

echo "---------------- 启动生产 ----------------"
docker-compose up -d --build prod-full

echo "---------------- 查看容器日志 ----------------"
docker logs -f gower
