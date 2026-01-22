@echo off
echo ========================================
echo PostgreSQL Setup Instructions
echo ========================================
echo.
echo You need PostgreSQL to run this application.
echo.
echo OPTION 1: Install PostgreSQL locally
echo 1. Download from: https://www.postgresql.org/download/windows/
echo 2. Install with default settings
echo 3. Remember the password you set for 'postgres' user
echo 4. Update .env file with your password
echo.
echo OPTION 2: Use Docker (Recommended for development)
echo Run this command in PowerShell:
echo.
echo docker run --name auth-postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=authdb -p 5432:5432 -d postgres:15
echo.
echo ========================================
echo After PostgreSQL is ready:
echo ========================================
echo 1. Update the .env file with correct credentials
echo 2. Run: go run main.go
echo.
pause
