#!/bin/bash

if ! command -v rclone &> /dev/null
then
    echo "---------------- rclone installing... ----------------"
    go install github.com/rclone/rclone@v1.62.2
    echo "---------------- rclone installed ----------------"
else
    echo "---------------- rclone is already installed ----------------"
fi

rclone version

echo "---------------- please create test and prod ssh server ----------------"
rclone config
