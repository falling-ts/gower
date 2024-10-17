#!/bin/bash

echo "---------------- 清理容器 ----------------"
docker-compose down

echo "---------------- 启动测试 ----------------"
docker-compose up -d --build test-full

echo "---------------- 查看容器日志 ----------------"
docker logs -f gower
