@echo off

echo ---------------- build static... ----------------
call npm run prod

echo ---------------- go test... ----------------
go test -tags prod,tmpl,static
REM go test -bench=Benchmark -tags prod,tmpl,static

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

echo ---------------- go build ----------------
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o gower -tags prod,tmpl,static

echo ---------------- uploading...----------------
rclone mkdir prod:/go/bin
rclone deletefile --progress prod:/go/bin/gower
rclone copy --progress ./ prod:/go/bin/ ^
    --include "envs/.env.development" ^
    --include "envs/.env.production" ^
    --include "public/static/**" ^
    --include "storage/**" ^
    --include "third_apps/**" ^
    --include "gower" ^
    --include "docker-compose.yaml" ^
    --include "Dockerfile" ^
    --include "run.sh"

echo ---------------- finished [next connect ssh and run] ----------------
