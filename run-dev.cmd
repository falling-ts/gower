@echo off

echo ---------------- build static... ----------------
call npm run dev

echo ---------------- start dev ----------------
docker compose up -d dev
