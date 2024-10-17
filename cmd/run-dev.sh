#!/bin/bash

echo "---------------- build static... ----------------"
npm run dev

echo "---------------- go test... ----------------"
go test -tags tmpl,static
# go test -bench=Benchmark -tags tmpl,static

echo "---------------- go build ----------------"
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go build -o gower -tags tmpl,static

echo "---------------- clean docker... ----------------"
docker-compose down

echo "---------------- start dev ----------------"
docker-compose up -d --build gower

echo "---------------- tail -f dev log ----------------"
docker logs -f gower
