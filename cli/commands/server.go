package commands

import (
	"github.com/lechgu/haproxylab/clients"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var serverListParams struct {
	baseURL  string
	username string
	password string
	backend  string
}

var serverRemoveParams struct {
	baseURL  string
	username string
	password string
	backend  string
	name     string
}

var serverAddParams struct {
	baseURL  string
	username string
	password string
	backend  string
	name     string
	address  string
	port     int
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Manage servers",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var serverListCommand = &cobra.Command{
	Use:   "list",
	Short: "List servers",
	RunE: func(cmd *cobra.Command, args []string) error {
		haClient := clients.New(serverListParams.baseURL, serverListParams.username, serverListParams.password)
		servers, err := haClient.ListServers(serverListParams.backend)
		if err != nil {
			return err
		}
		tbl := table.New("Name", "Address", "Port")
		for _, server := range servers {
			tbl.AddRow(server.Name, server.Address, server.Port)
		}
		tbl.Print()
		return nil
	},
}

func init() {
	serverListCommand.Flags().StringVarP(&serverListParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	serverListCommand.MarkFlagRequired("base-url")
	serverListCommand.Flags().StringVarP(&serverListParams.username, "username", "u", "", "Dataplane API user name")
	serverListCommand.MarkFlagRequired("username")
	serverListCommand.Flags().StringVarP(&serverListParams.password, "password", "p", "", "Dataplane API user password")
	serverListCommand.MarkFlagRequired("password")
	serverListCommand.Flags().StringVarP(&serverListParams.backend, "backend", "n", "", "Backend name")
	serverListCommand.MarkFlagRequired("backend")
	serverCmd.AddCommand(serverListCommand)
	rootCmd.AddCommand(serverCmd)
}
