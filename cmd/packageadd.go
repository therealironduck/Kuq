package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var packageAddCmd = &cobra.Command{
	Use:   "package:add",
	Short: "Add a new package to Kuq",
	Long:  `Add a new package to Kuq. Leave arguments empty for interactive screen`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Add new package")
	},
}
