package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func myLs(args []string) int {
	dirs := args
	if len(dirs) == 0 {
		dirs = []string{"."}
	}
	fmt.Printf("current directory %s \n", cwd)
	for _, sub := range dirs {
		tmp := filepath.Join(cwd, sub)
		d, err := root.Open(tmp)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ls:", err)
			return 1
		}
		entries, err := d.ReadDir(-1)
		d.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "ls:", err)
			return 1
		}
		for _, e := range entries {
			name := e.Name()
			if strings.HasPrefix(name, ".") {
				continue
			}
			if e.IsDir() {
				name += "/"
			}
			info, err := e.Info()
			if err != nil {
				continue
			}
			fmt.Printf("%s\t%d\t%s\n", name, info.Size(), info.ModTime().Format("2006-01-02 15:04:05"))
		}
	}
	return 0
}
