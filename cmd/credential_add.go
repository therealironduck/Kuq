package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therealironduck/kuq/internal/credential"
	"github.com/therealironduck/kuq/internal/global"
)

var credentialAddCmd = &cobra.Command{
	Use:   "credential:add [name] [sshKey]",
	Short: "Add a new ssh key to Kuq",
	Long:  `Add a new ssh key to Kuq. Leave arguments empty for interactive screen`,
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			// TODO: TUI
			fmt.Printf("TODO: Integrate interactive TUI")
			return
		}

		name := args[0]
		sshKey := args[1]

		app, _ := global.NewApp(flagWorkspace, nil) // TODO: Handle error

		id, _ := credential.Add(cmd.Context(), app, credential.AddOptions{Name: name, SSHKey: sshKey}) // TODO: Error handling

		fmt.Printf("New credential created with ID #%d\n", id) // TODO: Nice formatting
	},
}
