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
	port     int32
}

var serverCommand = &cobra.Command{
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

var serverRemoveCommand = &cobra.Command{
	Use:   "remove",
	Short: "Remove a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		haClient := clients.New(serverRemoveParams.baseURL, serverRemoveParams.username, serverRemoveParams.password)
		return haClient.RemoveServer(serverRemoveParams.name, serverRemoveParams.backend)
	},
}

var serverAddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a server",
	RunE: func(cmd *cobra.Command, args []string) error {
		haClient := clients.New(serverAddParams.baseURL, serverAddParams.username, serverAddParams.password)
		return haClient.AddServer(serverAddParams.name, serverAddParams.address, int(serverAddParams.port), serverAddParams.backend)
	},
}

func init() {
	serverListCommand.Flags().StringVarP(&serverListParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	serverListCommand.MarkFlagRequired("base-url")
	serverListCommand.Flags().StringVarP(&serverListParams.username, "username", "u", "", "Dataplane API user name")
	serverListCommand.MarkFlagRequired("username")
	serverListCommand.Flags().StringVarP(&serverListParams.password, "password", "p", "", "Dataplane API user password")
	serverListCommand.MarkFlagRequired("password")
	serverListCommand.Flags().StringVarP(&serverListParams.backend, "backend", "e", "", "Backend name")
	serverListCommand.MarkFlagRequired("backend")
	serverCommand.AddCommand(serverListCommand)

	serverRemoveCommand.Flags().StringVarP(&serverRemoveParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	serverRemoveCommand.MarkFlagRequired("base-url")
	serverRemoveCommand.Flags().StringVarP(&serverRemoveParams.username, "username", "u", "", "Dataplane API user name")
	serverRemoveCommand.MarkFlagRequired("username")
	serverRemoveCommand.Flags().StringVarP(&serverRemoveParams.password, "password", "p", "", "Dataplane API user password")
	serverRemoveCommand.MarkFlagRequired("password")
	serverRemoveCommand.Flags().StringVarP(&serverRemoveParams.backend, "backend", "e", "", "Backend name")
	serverRemoveCommand.MarkFlagRequired("backend")
	serverRemoveCommand.Flags().StringVarP(&serverRemoveParams.name, "name", "n", "", "Server name")
	serverRemoveCommand.MarkFlagRequired("name")
	serverCommand.AddCommand(serverRemoveCommand)

	serverAddCommand.Flags().StringVarP(&serverAddParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	serverAddCommand.MarkFlagRequired("base-url")
	serverAddCommand.Flags().StringVarP(&serverAddParams.username, "username", "u", "", "Dataplane API user name")
	serverAddCommand.MarkFlagRequired("username")
	serverAddCommand.Flags().StringVarP(&serverAddParams.password, "password", "p", "", "Dataplane API user password")
	serverAddCommand.MarkFlagRequired("password")
	serverAddCommand.Flags().StringVarP(&serverAddParams.backend, "backend", "e", "", "Backend name")
	serverAddCommand.MarkFlagRequired("backend")
	serverAddCommand.Flags().StringVarP(&serverAddParams.name, "name", "n", "", "Server name")
	serverAddCommand.MarkFlagRequired("name")
	serverAddCommand.Flags().StringVarP(&serverAddParams.address, "address", "a", "", "Server address")
	serverAddCommand.MarkFlagRequired("address")
	serverAddCommand.Flags().Int32Var(&serverAddParams.port, "port", 80, "Server port")
	serverAddCommand.MarkFlagRequired("port")
	serverCommand.AddCommand(serverAddCommand)

	rootCmd.AddCommand(serverCommand)
}
