package main

import (
	"github.com/seadiaz/adoption/src/client"
	"github.com/spf13/cobra"
)

const defaultURL = "http://localhost:3000"

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "adoption",
		Short: "A generator for Cobra based Applications",
		Long:  "adoption cli for interacting with adoption server",
	}

	loadCmd = &cobra.Command{
		Use:   "load <kind>",
		Short: "load csv data into adoption server",
		Long:  "load csv data into adoption server",
		Args:  cobra.ExactArgs(1),
		Run:   doLoadData,
	}
)

func mainCLI() {
	rootCmd.Execute()
}

func doLoadData(cmd *cobra.Command, args []string) {
	params := &client.Params{
		Filename: cmd.Flag("file").Value.String(),
		URL:      cmd.Flag("url").Value.String(),
		APIKey:   cmd.Flag("api-key").Value.String(),
		Kind:     args[0],
	}
	client.LoadData(params)
}

func init() {
	loadCmd.Flags().StringP("url", "u", defaultURL, "The URL of the running instance of adoption server")
	loadCmd.Flags().StringP("api-key", "k", "", "API Key which is going to be send by Authorization header")
	loadCmd.Flags().StringP("file", "f", "", "Load data from `FILE` (required)")
	loadCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(loadCmd)
}
