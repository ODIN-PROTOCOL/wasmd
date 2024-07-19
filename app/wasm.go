package app

import (
	wasmkeeper "github.com/ODIN-PROTOCOL/wasmd/x/wasm/keeper"
)

// Deprecated: Use BuiltInCapabilities from github.com/ODIN-PROTOCOL/wasmd/x/wasm/keeper
func AllCapabilities() []string {
	return wasmkeeper.BuiltInCapabilities()
}
