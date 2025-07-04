# go-quickjs-wasi

A Go module that embeds the QuickJS-NG WASI WebAssembly runtime.

## About QuickJS-NG

QuickJS is a small and embeddable JavaScript engine. It aims to support the latest ECMAScript specification.

This project uses [QuickJS-NG] which is a fork of the original [QuickJS project]
by Fabrice Bellard and Charlie Gordon, after it went dormant, with the intent of
reigniting its development.

[QuickJS-NG]: https://github.com/quickjs-ng/quickjs
[QuickJS project]: https://bellard.org/quickjs/

## Purpose

This module provides easy access to the QuickJS-NG JavaScript engine compiled to
WebAssembly with WASI support. The WASM binary is embedded directly in the Go
module, making it easy to use QuickJS in Go applications without external
dependencies.

## Features

- Embeds the latest QuickJS-NG WASI WebAssembly binary
- Provides version information about the embedded QuickJS release
- Automatic update script to fetch the latest QuickJS-NG release

## Usage

```go
package main

import (
    "fmt"
    "github.com/paralin/go-quickjs-wasi"
)

func main() {
    // Access the embedded WASM binary
    wasmBytes := quickjswasi.QuickJSWASM
    fmt.Printf("QuickJS WASM size: %d bytes\n", len(wasmBytes))

    // Get version information
    fmt.Printf("QuickJS version: %s\n", quickjswasi.Version)
    fmt.Printf("Download URL: %s\n", quickjswasi.DownloadURL)
}
```

## Updating

To update to the latest QuickJS-NG release:

```bash
./update-quickjs.bash
```

This script will:
1. Fetch the latest release information from the QuickJS-NG GitHub repository
2. Download the `qjs-wasi.wasm` file
3. Generate version information constants

## Testing

```bash
go test
```

## License

This module is released under the same license as the embedded QuickJS-NG project.

MIT
