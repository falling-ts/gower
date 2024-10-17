@echo off

echo ---------------- build static... ----------------
call npm run dev

echo ---------------- go test... ----------------
go test -tags tmpl,static
REM go test -bench=Benchmark -tags tmpl,static

echo ---------------- go build ----------------
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go build -o gower -tags tmpl,static

echo ---------------- clean docker... ----------------
docker-compose down

echo ---------------- start dev ----------------
docker-compose up -d --build gower

echo ---------------- tail -f dev log ----------------
docker logs -f gower
