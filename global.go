package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	root *os.Root
	cwd  string
)

func init() {
	exe, err := os.Executable()
	cwd = "."
	if err != nil {
		panic(err)
	}
	base := filepath.Dir(exe)
	fsPath := filepath.Join(base, "_fs")

	if _, err := os.Stat(fsPath); os.IsNotExist(err) {
		if wd, err := os.Getwd(); err == nil {
			fsPath = filepath.Join(wd, "_fs")
		}
	}
	r, err := os.OpenRoot(fsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open sandbox root:", err)
		os.Exit(1)
	}
	root = r
}
