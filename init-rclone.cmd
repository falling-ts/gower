@echo off

@echo "---------------- rclone installing... ----------------"
go install github.com/rclone/rclone@v1.62.2
@echo "---------------- rclone installed ----------------"

rclone version

@echo "---------------- please create test and prod ssh server ----------------"
rclone config
