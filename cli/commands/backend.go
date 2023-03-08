package commands

import (
	"github.com/lechgu/haproxylab/clients"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var backendListParams struct {
	baseURL  string
	username string
	password string
}

var backendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Manage backends",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var backendListCommand = &cobra.Command{
	Use:   "list",
	Short: "List backends",
	RunE: func(cmd *cobra.Command, args []string) error {
		haClient := clients.New(backendListParams.baseURL, backendListParams.username, backendListParams.password)
		backends, err := haClient.ListBackends()
		if err != nil {
			return err
		}
		tbl := table.New("Name", "Mode", "Connect Timeout", "Server Timeout")
		for _, backend := range backends {
			tbl.AddRow(backend.Name, backend.Mode, backend.ConnectTimeout, backend.ServerTimeout)
		}
		tbl.Print()

		return nil
	},
}

func init() {
	backendListCommand.Flags().StringVarP(&backendListParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	backendListCommand.MarkFlagRequired("base-url")
	backendListCommand.Flags().StringVarP(&backendListParams.username, "username", "u", "", "Dataplane API user name")
	backendListCommand.MarkFlagRequired("username")
	backendListCommand.Flags().StringVarP(&backendListParams.password, "password", "p", "", "Dataplane API user password")
	backendListCommand.MarkFlagRequired("password")

	backendCmd.AddCommand(backendListCommand)
	rootCmd.AddCommand(backendCmd)
}
