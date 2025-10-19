#!/bin/bash

# Build the binary
echo "Building slfy..."
if ! go build -o slfy main.go; then
    echo "Error: Failed to build slfy"
    exit 1
fi

# Make it executable
chmod +x slfy

# Move to /usr/local/bin
echo "Installing slfy to /usr/local/bin..."
if ! sudo mv slfy /usr/local/bin/; then
    echo "Error: Failed to install slfy. You may need to run with sudo or check permissions."
    exit 1
fi

echo "✅ slfy installed successfully to /usr/local/bin/slfy"

# Check if /usr/local/bin is in PATH
if ! echo "$PATH" | grep -q "/usr/local/bin"; then
    echo ""
    echo "⚠️  WARNING: /usr/local/bin is not in your PATH"
    echo "Add this line to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "export PATH=\"/usr/local/bin:\$PATH\""
    echo ""
    echo "Then run: source ~/.zshrc  # or ~/.bashrc"
    echo ""
else
    echo "✅ /usr/local/bin is already in your PATH"
fi

echo ""
echo "You can now run 'slfy' from anywhere!"
