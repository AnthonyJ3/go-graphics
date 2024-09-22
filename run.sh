#!/bin/bash

# Exit immediately if a command exits with a non-zero status
set -e

# Build the Go project
echo "Building..."
go build -o go-graphics

# Run the built executable
echo "Running..."
./go-graphics
