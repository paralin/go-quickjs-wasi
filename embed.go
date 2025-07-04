package quickjswasi

import _ "embed"

// QuickJSWASM contains the binary contents of qjs-wasi.wasm.
//
// See: https://github.com/quickjs-ng/quickjs/releases/download/v0.10.1/qjs-wasi.wasm
//
//go:embed qjs-wasi.wasm
var QuickJSWASM []byte
