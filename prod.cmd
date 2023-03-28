@echo off

echo ---------------- start upload... ----------------
rclone mkdir test:/go/src
rclone copy ./ test:/go/src --exclude node_modules/** --exclude vendor/** --exclude .env.production
echo ---------------- finished ----------------
