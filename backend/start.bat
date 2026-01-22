@echo off
echo ====================================
echo Auth Dashboard Backend - Quick Start
echo ====================================
echo.

echo Checking Go installation...
go version
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Go is not installed or not in PATH
    pause
    exit /b 1
)
echo.

echo Installing dependencies...
go mod download
if %ERRORLEVEL% NEQ 0 (
    echo ERROR: Failed to download dependencies
    pause
    exit /b 1
)
echo.

echo Running database migrations and starting server...
echo Server will be available at http://localhost:8080
echo.
echo Press Ctrl+C to stop the server
echo.

go run main.go
