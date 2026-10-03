package main

import (
	"fmt"
	"os"
	"os/exec"
)

type builtin func(args []string) int

var builtins = map[string]builtin{
	"ls":    func(args []string) int { return run(append([]string{"ls", "--cwd=" + cwd}, args...)) },
	"mkdir": func(args []string) int { return run(append([]string{"mkdir"}, args...)) },
	"exit":  func(_ []string) int { os.Exit(0); return 0 },
	"quit":  func(_ []string) int { os.Exit(0); return 0 },
	"pwd":   func(args []string) int { fmt.Printf("%s\n", cwd); return 0 },
	"cd":    func(args []string) int { return chdir(args) },
}

var implementations = map[string]builtin{
	"__ls":    myLs,
	"__mkdir": mkDir,
}

func run(arg []string) int {
	if len(arg) < 1 {
		fmt.Println("Usage: cmd [args...]")
		return 0
	}
	self, err := os.Executable()
	fieldName, args := arg[0], arg[1:]
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("unable to execute %s", fieldName), err)
		return 1
	}
	tmp := fmt.Sprintf("__%s", fieldName)
	cmd := exec.Command(self, append([]string{tmp}, args...)...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}
		fmt.Fprintln(os.Stderr, ":", err)
		return 1
	}
	return 0
}
