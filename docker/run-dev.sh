#!/bin/bash

echo "# npm run dev"
npm run dev
echo

echo "# go test -tags tmpl,static"
go test -tags tmpl,static
# go test -bench=Benchmark -tags tmpl,static
echo

echo "# go build -o gower -tags tmpl,static"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -o gower -tags tmpl,static
echo

echo "# docker compose down"
docker compose down
echo

echo "# docker compose up -d --build gower"
docker compose up -d --build gower
echo

echo "# docker logs -f gower"
docker logs -f gower
