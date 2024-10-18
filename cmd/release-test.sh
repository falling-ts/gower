#!/bin/bash

echo "---------------- build static... ----------------"
npm run test

echo "---------------- go test [need edit envs/.env.test DB_DRIVER as sqlite]... ----------------"
go test -tags test,tmpl,static
# go test -bench=Benchmark -tags test,tmpl,static

echo "---------------- go build ----------------"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -o gower -tags test,tmpl,static

echo "---------------- uploading...----------------"
rclone mkdir test:go/bin
rclone deletefile --progress test:go/bin/gower
rclone copy --progress ./ test:go/bin/ \
    --include "gower" \
    --include "cmd/run.sh" \
    --include "gower.service"

echo "---------------- finished [next connect ssh and run] ----------------"
