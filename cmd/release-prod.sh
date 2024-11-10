@echo off

echo "# npm run prod"
npm run prod
echo

echo "# go test -tags prod,tmpl,static"
echo "[Notice]: need edit envs/.env.prod DB_DRIVER as sqlite"
go test -tags prod,tmpl,static
# go test -bench=Benchmark -tags prod,tmpl,static
echo

echo "# go build -o gower -tags prod,tmpl,static"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o gower -tags prod,tmpl,static
echo

echo "# rclone mkdir gower-prod:/go/bin/gower"
rclone mkdir gower-prod:/go/bin/gower
echo

echo "# rclone deletefile --progress gower-prod:/go/bin/gower/gower"
rclone deletefile --progress gower-prod:/go/bin/gower/gower
echo

echo "# rclone copy --progress ./ gower-prod:/go/bin/gower/ --include ..."
rclone copy --progress ./ gower-prod:/go/bin/gower/ \
    --include "gower" \
    --include "cmd/run.sh" \
    --include "gower.service"
echo

echo "[Notice]: next connect ssh and run"
