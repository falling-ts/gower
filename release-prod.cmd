@echo off

echo ---------------- build static... ----------------
call npm run prod


echo ---------------- clean temp... ----------------
del /s /q *.log
del /s /q *.db
del /s /q *.cache

cd third_apps/tidb

del /s /q data
rmdir /s /q data
mkdir data
echo *> data\.gitignore
echo !.gitignore>> data\.gitignore

del /s /q logs
rmdir /s /q logs
mkdir logs
echo *> logs/.gitignore
echo !.gitignore>> logs\.gitignore

cd ../mysql5.7

del /s /q data
rmdir /s /q data
mkdir data
echo *> data/.gitignore
echo !.gitignore>> data\.gitignore

cd ../../storage/app

del /s /q uploads
rmdir /s /q uploads
mkdir uploads

cd ../../

echo ---------------- uploading... ----------------
rclone mkdir prod:/go/src
rclone copy --progress ^
 ./ prod:/go/src ^
 --exclude .git/** ^
 --exclude .github/** ^
 --exclude .idea/** ^
 --exclude node_modules/** ^
 --exclude vendor/** ^
 --exclude .env.test ^
 --exclude Dockerfile-development ^
 --exclude Dockerfile-test ^
 --exclude run-dev.cmd ^
 --exclude run-test.sh ^
 --exclude dev-entrypoint.sh ^
 --exclude test-entrypoint.sh

echo ---------------- finished ----------------

