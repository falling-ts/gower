@echo off

echo ---------------- build static... ----------------
call npm run test


echo ---------------- clean temp... ----------------
del /s /q *.log
del /s /q *.db

cd third_apps/tidb

del /s /q data
rd /s /q data
mkdir data
echo *> data\.gitignore
echo !.gitignore>> data\.gitignore

del /s /q logs
rd /s /q logs
mkdir logs
echo *> logs/.gitignore
echo !.gitignore>> logs\.gitignore

cd ../mysql5.7

del /s /q data
rd /s /q data
mkdir data
echo *> data/.gitignore
echo !.gitignore>> data\.gitignore

cd ../../../

echo ---------------- uploading... ----------------
rclone mkdir test:/go/src
rclone copy ^
 ./ test:/go/src ^
 --exclude node_modules/** ^
 --exclude vendor/** ^
 --exclude .env.production ^
 --exclude Dockerfile-development ^
 --exclude Dockerfile-production ^
 --exclude run-prod.sh

echo ---------------- finished ----------------

