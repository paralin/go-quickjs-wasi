#!/bin/bash

set -euo pipefail

# Get the latest release info from GitHub API
echo "Fetching latest QuickJS-NG release information..."
RELEASE_INFO=$(curl -s "https://api.github.com/repos/quickjs-ng/quickjs/releases/latest")

# Extract version and download URL
VERSION=$(echo "$RELEASE_INFO" | grep '"tag_name"' | cut -d'"' -f4)
DOWNLOAD_URL=$(echo "$RELEASE_INFO" | grep '"browser_download_url"' | grep 'qjs-wasi.wasm' | cut -d'"' -f4)

if [ -z "$VERSION" ] || [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not find version or download URL"
    exit 1
fi

echo "Latest version: $VERSION"
echo "Download URL: $DOWNLOAD_URL"

# Download the WASM file
echo "Downloading qjs-wasi.wasm..."
curl -L -o qjs-wasi.wasm "$DOWNLOAD_URL"

echo "Downloaded qjs-wasi.wasm successfully"

# Generate version info Go file
echo "Generating version.go..."
cat > version.go << EOF
package quickjswasi

// QuickJS-NG version information
const (
	// Version is the QuickJS-NG version
	Version = "$VERSION"
	// DownloadURL is the URL where this WASM file was downloaded from
	DownloadURL = "$DOWNLOAD_URL"
)
EOF

echo "Generated version.go with version $VERSION"
echo "Update complete!"
