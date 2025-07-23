
#!/bin/bash

# AdminiSoftware Test Runner
# Runs all backend tests

set -e

echo "AdminiSoftware Test Suite"
echo "========================"

# Set test environment
export GO_ENV=test
export GIN_MODE=test

# Navigate to backend directory
cd "$(dirname "$0")/.."

echo "Running Go tests..."

# Run unit tests
echo "→ Running unit tests..."
go test ./internal/... -v -cover

# Run integration tests
echo "→ Running integration tests..."
go test ./tests/integration/... -v

# Run API tests
echo "→ Running API tests..."
go test ./tests/api/... -v

# Generate coverage report
echo "→ Generating coverage report..."
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html

echo "✓ All tests completed successfully!"
echo "Coverage report generated: coverage.html"
