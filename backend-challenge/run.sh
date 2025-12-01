#!/usr/bin/env bash
set -e

# Move to script directory (backend-challenge)
cd "$(dirname "$0")"

echo "🔧 Running go mod tidy..."
go mod tidy

echo "🚀 Starting Backend Server..."
go run cmd/server/main.go
