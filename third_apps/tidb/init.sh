#!/bin/sh

# 修改 root 密码
mysql -h 127.0.0.1 -P 4000 -u root < /init.sql
echo "MySQL root password has been updated."

pkill -f tidb-server
