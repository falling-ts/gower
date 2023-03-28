@echo off

echo ---------------- build static... ----------------
call npm run test

echo ---------------- go test [need edit envs/.env.test DB_DRIVER as sqlite]... ----------------
go test -bench=Benchmark -tags test,env,tmpl,static

echo ---------------- go build ----------------
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o output/gower -tags test,env,tmpl,static

echo ---------------- uploading...----------------
rclone mkdir test:/go/bin
rclone copyto --progress output/gower test:/go/bin/gower

echo ---------------- finished [next connect ssh and run] ----------------
