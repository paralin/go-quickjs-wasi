package main

import (
	"context"
	"embed"
	"os"
	"testing"

	quickjswasi "github.com/paralin/go-quickjs-wasi"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

//go:embed repl_test.js
var testFS embed.FS

func TestRunJavaScriptFile(t *testing.T) {
	// Create a new WebAssembly Runtime
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)

	// Configure the module with stdin, stdout, stderr, and embedded filesystem
	config := wazero.NewModuleConfig().
		WithStdin(os.Stdin).
		WithStdout(os.Stdout).
		WithStderr(os.Stderr).
		WithFS(testFS)

	// Instantiate WASI
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the Wasm module with the JS file as argument
	args := []string{quickjswasi.QuickJSWASMFilename, "--std", "repl_test.js"}
	mod, err := r.InstantiateWithConfig(ctx, quickjswasi.QuickJSWASM, config.WithArgs(args...))
	if err != nil {
		// Check if it's an exit error with code 0 (successful exit)
		if exitErr, ok := err.(*sys.ExitError); ok {
			if exitErr.ExitCode() != 0 {
				t.Fatalf("QuickJS exited with non-zero code: %d", exitErr.ExitCode())
			}
			// Exit code 0 is success
		} else {
			t.Fatalf("Failed to instantiate module: %v", err)
		}
	}
	_ = mod

	t.Log("Successfully executed JavaScript file")
}
