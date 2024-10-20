@echo off

echo # npm run test
call npm run test
echo.

echo # clean
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
echo.

echo # rclone mkdir gower-test:/go/src/gower
rclone mkdir gower-test:/go/src/gower
echo.

echo # rclone copy --progress ./ gower-test:/go/src/gower/ --include ...
rclone copy --progress ./ gower-test:/go/src/gower/ ^
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
    --include "go.mod" ^
    --include "go.sum" ^
    --include "main.go" ^
    --include "main_test.go" ^
    --include "docker/Dockerfile-test-full" ^
    --include "docker/entrypoint-test-full.sh" ^
    --include "docker/run-test-full.sh"
echo.

echo [Notice]: next connect ssh and run
