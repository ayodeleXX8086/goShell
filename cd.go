package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func chdir(args []string) int {
	target := "."
	if len(args) > 0 {
		target = args[0]
	}
	tmp := filepath.Join(cwd, target)
	next := filepath.Clean(tmp)
	if next == ".." || strings.HasPrefix(next, "../") {
		fmt.Fprintln(os.Stderr, "cd: cannot escape sandbox root")
		return 1
	}
	if next == "." {
		cwd = "."
		return 0
	}
	d, err := root.Open(next)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cd: ", err)
		return 1
	}
	info, err := d.Stat()
	d.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cd: ", err)
		return 1
	}
	if !info.IsDir() {
		fmt.Fprintln(os.Stderr, "cd: not a directory")
		return 1
	}
	cwd = next
	return 0
}
