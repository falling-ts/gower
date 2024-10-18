@echo off

echo "---------------- build static... ----------------"
npm run prod

echo "---------------- go test... ----------------"
# go test -tags prod,tmpl,static
# go test -bench=Benchmark -tags prod,tmpl,static

echo "---------------- clean temp... ----------------"
rm -rf ./*.log
rm -rf ./*.db
rm -rf ./*.cache

cd ./third_apps/tidb || exit

find data ! -path data/.gitignore -exec rm -rf {} \;
find logs ! -path logs/.gitignore -exec rm -rf {} \;

cd ../mysql5.7 || exit

find data ! -path data/.gitignore -exec rm -rf {} \;

cd ../../storage/app || exit

rm -rf upload/*

cd ../../

echo "---------------- go build ----------------"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -o gower -tags prod,tmpl,static

echo "---------------- uploading...----------------"
rclone mkdir prod:go/bin
rclone deletefile --progress prod:go/bin/gower
rclone copy --progress ./ prod:go/bin/ \
    --include "envs/.env.development" \
    --include "envs/.env.production" \
    --include "public/static/**" \
    --include "storage/**" \
    --include "third_apps/**" \
    --include "gower" \
    --include "docker-compose.yaml" \
    --include "docker/Dockerfile" \
    --include "docker/run.sh"

echo "---------------- finished [next connect ssh and run] ----------------"
