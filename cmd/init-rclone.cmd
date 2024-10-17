@echo off

where rclone >nul 2>&1
if %errorlevel% == 0 (
    echo ---------------- rclone installed ----------------
) else (
    echo ---------------- rclone installing... ----------------
    go install github.com/rclone/rclone@v1.62.2
    echo ---------------- rclone installed ----------------
)


rclone version

echo ---------------- please create test and prod ssh server ----------------
rclone config
