package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		if root == nil {
			return
		}
		err := root.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot close sandbox root:", err)
			os.Exit(1)
		}
	}()
	if len(os.Args) >= 2 {
		if fn, ok := implementations[os.Args[1]]; ok {
			fs := flag.NewFlagSet(os.Args[2], flag.ContinueOnError)
			cwdFlag := fs.String("cwd", ".", "virtual cwd inside sandbox")
			if err := fs.Parse(os.Args[2:]); err != nil {
				os.Exit(2)
			}
			cwd = *cwdFlag
			os.Exit(fn(os.Args[3:]))
		} else {
			fmt.Fprintln(os.Stderr, "unknown implementation:", os.Args[1])
			os.Exit(1)
		}
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(text)
		if len(fields) == 0 {
			fmt.Print("> ")
			continue
		}
		name, args := fields[0], fields[1:]
		if fn, ok := builtins[name]; ok {
			fn(args)
		} else {
			fmt.Printf("got %s\n", text)
		}
		fmt.Print("> ")
	}
}
