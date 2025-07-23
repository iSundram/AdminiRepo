
#!/bin/bash

# AdminiSoftware Database Migration Script
# Handles database schema migrations

set -e

echo "AdminiSoftware Database Migration"
echo "================================"

# Load environment variables
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Navigate to backend directory
cd "$(dirname "$0")/.."

echo "Starting database migration..."

# Run migration command
go run cmd/migrate.go

echo "✓ Database migration completed successfully!"
