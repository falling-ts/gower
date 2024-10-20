#!/bin/bash

echo "# npm run test "
npm run test
echo

echo "# go test -tags test,tmpl,static"
echo "[Notice]: need edit envs/.env.test DB_DRIVER as sqlite"
go test -tags test,tmpl,static
# go test -bench=Benchmark -tags test,tmpl,static
echo

echo "# go build -o gower -tags test,tmpl,static"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o gower -tags test,tmpl,static
echo

echo "# rclone mkdir gower-test:/go/bin/gower"
rclone mkdir gower-test:/go/bin/gower
echo

echo "# rclone deletefile --progress gower-test:/go/bin/gower/gower"
rclone deletefile --progress gower-test:/go/bin/gower/gower
echo

echo "# rclone copy --progress ./ gower-test:/go/bin/gower/ --include ..."
rclone copy --progress ./ gower-test:/go/bin/gower/ \
    --include "gower" \
    --include "cmd/run.sh" \
    --include "gower.service"
echo

echo "[Notice]: next connect ssh and run"
