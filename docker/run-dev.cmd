@echo off

echo # npm run dev
call npm run dev
echo.

echo # go test -tags tmpl,static
go test -tags tmpl,static
REM go test -bench=Benchmark -tags tmpl,static
echo.

echo # go build -o gower -tags tmpl,static
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o gower -tags tmpl,static
echo.

echo # docker compose down
docker compose down
echo.

echo # docker compose up -d --build gower
docker compose up -d --build gower
echo.

echo # docker logs -f gower
docker logs -f gower
