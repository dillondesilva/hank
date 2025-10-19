#!/bin/bash

# Build the binary
echo "Building hank..."
if ! go build -o hank main.go; then
    echo "Error: Failed to build hank"
    exit 1
fi

# Make it executable
chmod +x hank

# Move to /usr/local/bin
echo "Installing hank to /usr/local/bin..."
if ! sudo mv hank /usr/local/bin/; then
    echo "Error: Failed to install hank. You may need to run with sudo or check permissions."
    exit 1
fi

echo "✅ hank installed successfully to /usr/local/bin/hank"

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
echo "You can now run 'hank' from anywhere!"
