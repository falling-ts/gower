#!/bin/bash

@echo off

echo "---------------- build static... ----------------"
npm run prod


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

rm -rf uploads/*

cd ../../

echo "---------------- uploading... ----------------"
rclone mkdir prod:/go/src
rclone copy --progress \
 ./ prod:/go/src \
 --exclude .git/** \
 --exclude .github/** \
 --exclude .idea/** \
 --exclude node_modules/** \
 --exclude vendor/** \
 --exclude Dockerfile-development \
 --exclude Dockerfile-test \
 --exclude run-dev.sh \
 --exclude run-test.sh \
 --exclude dev-entrypoint.sh \
 --exclude test-entrypoint.sh

echo "---------------- finished [next connect ssh and run] ----------------"


