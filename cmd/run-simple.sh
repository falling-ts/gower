#!/bin/bash

echo "---------------- 卸载旧服务 ----------------"
systemctl stop gower
systemctl disable gower
rm -rf /etc/systemd/system/gower.service
systemctl daemon-reload
systemctl reset-failed

echo "---------------- 安装服务 ----------------"
chmod +x gower
cp gower.service /etc/systemd/system/


echo "---------------- 启动进程 ----------------"
systemctl daemon-reload
systemctl restart gower
systemctl enable gower
