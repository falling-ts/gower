@echo off

echo ---------------- build static... ----------------
call npm run prod

echo ---------------- go test [need edit envs/.env.production DB_DRIVER as sqlite]... ----------------
go test -tags prod,tmpl,static
REM go test -bench=Benchmark -tags prod,tmpl,static

echo ---------------- go build ----------------
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o gower -tags prod,tmpl,static

echo ---------------- uploading...----------------
rclone mkdir prod:go/bin
rclone deletefile --progress prod:go/bin/gower
rclone copy --progress ./ prod:go/bin/ ^
    --include "gower" ^
    --include "cmd/run.sh" ^
    --include "gower.service"

echo ---------------- finished [next connect ssh and run] ----------------
