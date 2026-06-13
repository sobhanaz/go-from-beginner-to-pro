// Session 15 — reading and writing files (and JSON to/from a file).
// Run:  go run examples/session15/files/files.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	AppName string `json:"app_name"`
	Port    int    `json:"port"`
	Debug   bool   `json:"debug"`
}

func main() {
	// Use a temp directory so we don't leave files lying around.
	path := filepath.Join(os.TempDir(), "golearn_config.json")

	// --- Write a whole file in one call (creates or overwrites) ---
	// 0644 is the file permission: owner read/write, others read.
	if err := os.WriteFile(path, []byte("hello, file!\n"), 0644); err != nil {
		fmt.Println("write error:", err)
		return
	}
	fmt.Println("wrote text to", path)

	// --- Read a whole file in one call ---
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}
	fmt.Printf("read back: %q\n", string(content))

	// --- Write JSON to a file ---
	cfg := Config{AppName: "TaskFlow", Port: 8080, Debug: true}
	jsonBytes, _ := json.MarshalIndent(cfg, "", "  ")
	if err := os.WriteFile(path, jsonBytes, 0644); err != nil {
		fmt.Println("write json error:", err)
		return
	}

	// --- Read JSON back into a struct ---
	raw, _ := os.ReadFile(path)
	var loaded Config
	if err := json.Unmarshal(raw, &loaded); err != nil {
		fmt.Println("parse json error:", err)
		return
	}
	fmt.Printf("loaded config: %+v\n", loaded)

	// Clean up the temp file.
	_ = os.Remove(path)
	fmt.Println("cleaned up temp file")
}
