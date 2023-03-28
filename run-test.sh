#!/bin/bash

echo "---------------- 启动测试 ----------------"
docker compose up -d test tidb grafana loki promtail

echo "---------------- 完毕 ----------------"
