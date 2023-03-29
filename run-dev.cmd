@echo off

echo ---------------- build static... ----------------
call npm run dev

echo ---------------- clean docker... ----------------
docker compose down

echo ---------------- start dev ----------------
docker compose up -d --build dev
