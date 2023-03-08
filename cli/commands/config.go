package commands

import (
	"fmt"

	"github.com/lechgu/haproxylab/clients"
	"github.com/spf13/cobra"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "Peek into the HAProxy configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configVersionParams struct {
	baseURL  string
	username string
	password string
}

var configVersionCommand = &cobra.Command{
	Use:   "version",
	Short: "get configuration version",
	RunE: func(cmd *cobra.Command, args []string) error {
		haClient := clients.New(configVersionParams.baseURL, configVersionParams.username, configVersionParams.password)
		ver, err := haClient.GetConfigurationVersion()
		if err != nil {
			return err
		}
		fmt.Printf("%d\n", ver)
		return nil
	},
}

func init() {
	configVersionCommand.Flags().StringVarP(&configVersionParams.baseURL, "base-url", "b", "", "Dataplane API base URL")
	configVersionCommand.MarkFlagRequired("base-url")
	configVersionCommand.Flags().StringVarP(&configVersionParams.username, "username", "u", "", "Dataplane API user name")
	configVersionCommand.MarkFlagRequired("username")
	configVersionCommand.Flags().StringVarP(&configVersionParams.password, "password", "p", "", "Dataplane API user password")
	configVersionCommand.MarkFlagRequired("password")
	configCommand.AddCommand(configVersionCommand)
	rootCmd.AddCommand(configCommand)
}
