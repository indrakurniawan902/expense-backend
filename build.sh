#!/bin/bash

# Build script untuk production

echo "🔨 Building Expense API..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.25.4 or higher."
    exit 1
fi

# Clean previous build
rm -f expense-api

# Download dependencies
echo "📦 Downloading dependencies..."
go mod download

# Build for Linux (untuk deploy ke server)
echo "🔨 Building binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o expense-api .

if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo "📊 Binary size: $(du -h expense-api | cut -f1)"
else
    echo "❌ Build failed!"
    exit 1
fi

# Build Docker image
echo "🐳 Building Docker image..."
docker build -t expense-api:latest .

if [ $? -eq 0 ]; then
    echo "✅ Docker image built successfully!"
    echo "🚀 To run: docker run -p 8080:8080 expense-api:latest"
else
    echo "❌ Docker build failed!"
    exit 1
fi
