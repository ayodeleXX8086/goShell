package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func mkDirAll(root *os.Root, path string, perm fs.FileMode) error {
	parts := strings.Split(path, string(os.PathSeparator))
	current := cwd
	for _, part := range parts {
		if part == "" {
			continue
		}
		if current == "" {
			current = part
		} else {
			current = filepath.Join(current, part)
		}
		err := root.Mkdir(current, perm)
		if err != nil {
			if os.IsExist(err) {
				continue
			}
			return err
		}
	}
	return nil
}

func mkDir(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: mkdir <dir>")
		return 1
	}
	err := mkDirAll(root, args[0], fs.FileMode(0755))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
