@echo off

echo # npm run test
call npm run test
echo.

echo # go test -tags test,tmpl,static
echo [Notice]: need edit envs/.env.test DB_DRIVER as sqlite
go test -tags test,tmpl,static
:: go test -bench=Benchmark -tags test,tmpl,static
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

echo # go build -o gower -tags test,tmpl,static
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o gower -tags test,tmpl,static
echo.

echo # rclone mkdir gower-test:/go/bin/gower
rclone mkdir gower-test:/go/bin/gower
echo.

echo # rclone deletefile --progress gower-test:/go/bin/gower/gower
rclone deletefile --progress gower-test:/go/bin/gower/gower
echo.

echo # rclone copy --progress ./ gower-test:/go/bin/gower/ --include ...
rclone copy --progress ./ gower-test:/go/bin/gower/ ^
    --include "envs/.env.dev" ^
    --include "envs/.env.test" ^
    --include "public/static/**" ^
    --include "storage/**" ^
    --include "third_apps/**" ^
    --include "gower" ^
    --include "docker-compose.yaml" ^
    --include "docker/Dockerfile" ^
    --include "docker/run.sh"
echo.

echo [Notice]: next connect ssh and run
