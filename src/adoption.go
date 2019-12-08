package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/seadiaz/adoption/src/client"
	"github.com/seadiaz/adoption/src/server"
	"github.com/spf13/cobra"
)

func init() {
	flag.Set("logtostderr", "true")
	flag.Parse()
}

func main() {
	mainCLI()
}

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

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "boot adoption server",
		Long:  "boot adoption server",
		Args:  cobra.ExactArgs(0),
		Run:   doBootServer,
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

func doBootServer(cmd *cobra.Command, args []string) {
	redisPort, _ := strconv.Atoi(cmd.Flag("redis-port").Value.String())
	params := &server.Params{
		Port:      cmd.Flag("port").Value.String(),
		Storage:   cmd.Flag("storage").Value.String(),
		RedisPort: redisPort,
		RedisHost: cmd.Flag("redis-host").Value.String(),
	}
	server.Boot(params)
}

func init() {
	loadCmd.Flags().StringP("url", "u", defaultURL, "The URL of the running instance of adoption server")
	loadCmd.Flags().StringP("api-key", "k", "", "API Key which is going to be send by Authorization header")
	loadCmd.Flags().StringP("file", "f", "", "Load data from `FILE` (required)")
	loadCmd.MarkFlagRequired("file")

	serverCmd.Flags().IntP("port", "p", 3000, "port the server will bind")
	serverCmd.Flags().StringP("storage", "s", "memory", "storage type where data going to be persisted")
	serverCmd.Flags().Int("redis-port", 6379, "redis port for using with redis storage")
	serverCmd.Flags().IP("redis-host", net.IPv4(127, 0, 0, 1), "redis host for using with redis storage")

	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(serverCmd)
}
