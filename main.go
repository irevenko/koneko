package main

import (
	"fmt"
	"os"

	cmd "github.com/irevenko/koneko/cmd"
)

func main() {
	cmd.AddCommands()

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
