package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	product = "todo"
	version = "0.0.1"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"ver"},
	Short:   fmt.Sprintf("Print the version of %s", product),
	Long:    fmt.Sprintf("All software has version. This is %s's", product),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s - v%s\n", product, version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
