#!/bin/bash

set -e

echo "=============================="
echo " PostgreSQL AUTH Setup Script"
echo "=============================="

# -----------------------
# USER INPUT
# -----------------------
read -p "Enter DB Name: " DB_NAME
read -p "Enter DB User: " DB_USER
read -p "Enter DB Host (default: localhost): " DB_HOST
read -p "Enter DB Port (default: 5432): " DB_PORT
read -s -p "Enter DB Password: " DB_PASSWORD
echo ""

# Defaults
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}

export PGPASSWORD=$DB_PASSWORD

echo ""
echo "Using:"
echo "DB_NAME=$DB_NAME"
echo "DB_USER=$DB_USER"
echo "DB_HOST=$DB_HOST"
echo "DB_PORT=$DB_PORT"
echo ""

# -----------------------
# CHECK DATABASE EXISTS
# -----------------------
echo "Checking if database exists..."

DB_EXISTS=$(psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -tAc "SELECT 1 FROM pg_database WHERE datname='$DB_NAME'")

if [ "$DB_EXISTS" = "1" ]; then
  echo "Database already exists ✅"
else
  echo "Creating database..."
  psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -c "CREATE DATABASE $DB_NAME;"
fi

# -----------------------
# APPLY SCHEMA
# -----------------------
echo "Applying schema..."

psql -U "$DB_USER" -h "$DB_HOST" -p "$DB_PORT" -d "$DB_NAME" -f auth_schema.sql

echo ""
echo "✅ Database setup completed successfully!"