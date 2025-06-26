#!/bin/bash

# Stop the server first
echo "Please stop the Go server (Ctrl+C) before running this script"
read -p "Press Enter when the server is stopped..."

# Backup current db
cp debt_tracker.db debt_tracker_backup.db
echo "Database backed up to debt_tracker_backup.db"

# remove the current db to recreate with new schema
rm debt_tracker.db
echo "Old DB removed"

# Start the server to recreate the database with new schema
echo "Now restart your Go server:"
echo "go run cmd/server/main.go"
echo ""
echo "The new database will allow multiple debts per contact!"
echo "You'll need to re-register your user and re-enter your data."
