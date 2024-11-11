#!/bin/bash

@echo off

echo "# npm run prod"
npm run prod
echo

echo "# clean"
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
echo

echo "# rclone mkdir gower-prod:/go/src/gower"
rclone mkdir gower-prod:/go/src/gower
echo

echo "# rclone copy --progress ./ gower-prod:/go/src/gower/ --include ..."
rclone copy --progress ./ gower-prod:/go/src/gower/ \
    --include "app/**" \
    --include "bootstrap/**" \
    --include "configs/**" \
    --include "envs/**" \
    --include "public/**" \
    --include "resources/**" \
    --include "routes/**" \
    --include "services/**" \
    --include "storage/**" \
    --include "tests/**" \
    --include "third_apps/**" \
    --include "trans/**" \
    --include "utils/**" \
    --include "docker-compose.yaml" \
    --include "go.mod" \
    --include "go.sum" \
    --include "main.go" \
    --include "main_test.go" \
    --include "docker/Dockerfile-prod-full" \
    --include "docker/entrypoint-prod-full.sh" \
    --include "docker/run-prod-full.sh"
echo

echo "[Notice]: next connect ssh and run"