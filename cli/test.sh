#!/bin/bash
#for creating new executables and deleted testing files

# Set filenames
EXECUTABLE="autodoc"
CONFIG_FILE="autodoc_config.json"
OUTPUT_FILE="output.txt"

# Remove existing files if they exist
echo "Removing old files..."
[ -f "$EXECUTABLE" ] && rm "$EXECUTABLE" && echo "Deleted $EXECUTABLE"
[ -f "$CONFIG_FILE" ] && rm "$CONFIG_FILE" && echo "Deleted $CONFIG_FILE"
[ -f "$OUTPUT_FILE" ] && rm "$OUTPUT_FILE" && echo "Deleted $OUTPUT_FILE"

# Build the Go executable
echo "Building new executable..."
go build -o "$EXECUTABLE"

echo "Done."

echo "running commands"
./autodoc health
./autodoc init
./autodoc generate > "output.txt"
