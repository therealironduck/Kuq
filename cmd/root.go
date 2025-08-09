package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const Version = "0.1.0"

var (
	flagWorkspace string
)

var rootCmd = &cobra.Command{
	Use:     "kuq",
	Short:   "Kuq is a fast, simple registry generator and server for Composer",
	Version: Version,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Hello World: %s\n", flagWorkspace)
	},
}

func init() {
	rootCmd.AddCommand(packageAddCmd)

	rootCmd.PersistentFlags().StringVarP(&flagWorkspace, "workspace", "w", ".", "The workspace where Kuq should store all data")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
