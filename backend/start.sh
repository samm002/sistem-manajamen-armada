#!/bin/sh

set -e

echo "Starting migration process..."

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be available..."
until pg_isready -h postgres -p 5432 -U $DB_USER; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

echo "PostgreSQL is available - checking database..."

export PGPASSWORD=$DB_PASSWORD

# Create database if it does not exist
psql -h postgres -U $DB_USER -tc "SELECT 1 FROM pg_database WHERE datname = '$DB_NAME'" | grep -q 1 || \
  psql -h postgres -U $DB_USER -c "CREATE DATABASE $DB_NAME;"

echo "Database created. Running migrations."

# Run database migrations
migrate -path /root/migrations/ -database $DB_URL up

# Check if migrations were successful
if [ $? -eq 0 ]; then
    echo "Migrations completed successfully"
else
    echo "Migration failed"
    exit 1
fi

echo "Starting main application..."
# Start the main application
exec ./main