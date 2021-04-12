package cmd

import (
	"fmt"

	tui "github.com/irevenko/koneko/tui"
	"github.com/spf13/cobra"
)

var Nyaa = &cobra.Command{
	Use:   "nyaa",
	Short: "Browse nyaa.si",
	Long:  `koneko nyaa`,
	Run: func(cmd *cobra.Command, args []string) {
		tui.Launch("nyaa")
	},
}

var Sukebei = &cobra.Command{
	Use:   "sukebei",
	Short: "Browse sukebei.nyaa.si",
	Long:  `koneko sukebei`,
	Run: func(cmd *cobra.Command, args []string) {
		tui.Launch("sukebei")
	},
}

var Help = &cobra.Command{
	Use:   "help",
	Short: "Help for the koneko",
	Long:  `koneko help`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(tui.HelpText)
	},
}

func AddCommands() {
	RootCmd.AddCommand(Nyaa)
	RootCmd.AddCommand(Sukebei)
	RootCmd.AddCommand(Help)
}
