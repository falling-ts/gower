@echo off

where rclone >nul 2>&1
if %errorlevel% neq 0 (
    echo # go install github.com/rclone/rclone@v1.62.2
    go install github.com/rclone/rclone@v1.62.2
    echo.
)

echo # rclone version
rclone version
echo.

echo # rclone config
echo [Notice]: please create gower-test and gower-prod ssh server
rclone config
