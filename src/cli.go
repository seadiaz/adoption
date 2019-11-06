package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/mitchellh/go-homedir"
	"github.com/seadiaz/adoption/src/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultURL = "http://localhost:3000"

var (
	// Used for flags.
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
		Run:   loadData,
	}
)

func mainCLI() {
	rootCmd.Execute()
}

func loadData(cmd *cobra.Command, args []string) {
	filename := cmd.Flag("file").Value.String()
	kind := args[0]
	client := &cli.Client{
		URL:      cmd.Flag("URL").Value.String(),
		Filename: filename,
	}
	switch kind {
	case "tools":
		client.LoadTools()
	case "people":
		client.LoadPeople()
	case "adoptions":
		client.LoadAdoptions()
	case "teams":
		client.LoadTeams()
	case "memberships":
		client.LoadMemberships()
	default:
		glog.Fatalf("kind %s not supported", kind)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("URL", "u", defaultURL, "The URL of the running instance of adoption server")

	loadCmd.Flags().StringP("file", "f", "", "Load data from `FILE` (required)")
	err := loadCmd.MarkFlagRequired("file")
	if err != nil {
		glog.Error(err)
	}

	rootCmd.AddCommand(loadCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			glog.Error(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
