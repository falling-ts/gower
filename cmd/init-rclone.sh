#!/bin/bash

if ! command -v rclone &> /dev/null
then
    echo "# go install github.com/rclone/rclone@v1.62.2"
    go install github.com/rclone/rclone@v1.62.2
    echo
fi

echo "# rclone version"
rclone version
echo

echo "# rclone config"
echo "[Notice]: please create gower-test and gower-prod ssh server"
rclone config
