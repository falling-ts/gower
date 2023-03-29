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

del /s /q upload
rmdir /s /q upload
mkdir upload

cd ../../

echo ---------------- uploading... ----------------
rclone mkdir prod:/go/src
rclone copy --progress ./ prod:/go/src/ ^
    --include "app/**" ^
    --include "bootstrap/**" ^
    --include "configs/**" ^
    --include "envs/**" ^
    --include "public/**" ^
    --include "resources/**" ^
    --include "routes/**" ^
    --include "services/**" ^
    --include "storage/**" ^
    --include "tests/**" ^
    --include "third_apps/**" ^
    --include "trans/**" ^
    --include "utils/**" ^
    --include "docker-compose.yaml" ^
    --include "Dockerfile-prod-full" ^
    --include "entrypoint-prod-full.sh" ^
    --include "go.mod" ^
    --include "go.sum" ^
    --include "main.go" ^
    --include "main_test.go" ^
    --include "run-prod-full.sh"

echo ---------------- finished [next connect ssh and run] ----------------

