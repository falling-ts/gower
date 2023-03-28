@echo off

echo ---------------- build static... ----------------
call npm run prod

echo ---------------- go test [need edit envs/.env.production DB_DRIVER as sqlite]... ----------------
go test -bench=Benchmark -tags prod,env,tmpl,static

echo ---------------- go build ----------------
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o output/gower -tags prod,env,tmpl,static

echo ---------------- uploading...----------------
rclone mkdir prod:/go/bin
rclone copyto --progress output/gower prod:/go/bin/gower

echo ---------------- finished [next connect ssh and run] ----------------
