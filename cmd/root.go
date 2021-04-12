package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd main cobra command
var RootCmd = &cobra.Command{
	Use:   "koneko",
	Short: "nyaa.si terminal client. Download East Asian torrents",
	Long:  `Complete documentation is available at https://github.com/irevenko/koneko`,
}
