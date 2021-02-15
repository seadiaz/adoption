package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/server"
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
		Run:   dispatchLoadCommand,
	}

	displayCmd = &cobra.Command{
		Use:   "display <kind>",
		Short: "display the information stored in the server",
		Long:  "display the information stored in the server",
		Args:  cobra.ExactArgs(1),
		Run:   dispatchDisplayCommand,
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

func dispatchLoadCommand(cmd *cobra.Command, args []string) {
	commandDispatcher := people.CreateCommandHandler(
		cmd.Flag("url").Value.String(),
		cmd.Flag("api-key").Value.String(),
	)
	cmd.Flags()
	params := &people.CommandHandlerParams{
		Filename: cmd.Flag("file").Value.String(),
		Action:   global.Load,
	}

	commandDispatcher.Execute(params)
}

func dispatchDisplayCommand(cmd *cobra.Command, args []string) {
	commandDispatcher := people.CreateCommandHandler(
		cmd.Flag("url").Value.String(),
		cmd.Flag("api-key").Value.String(),
	)
	cmd.Flags()
	params := &people.CommandHandlerParams{
		Action: global.Display,
	}

	commandDispatcher.Execute(params)
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

	displayCmd.Flags().StringP("url", "u", defaultURL, "The URL of the running instance of adoption server")
	displayCmd.Flags().StringP("api-key", "k", "", "API Key which is going to be send by Authorization header")

	serverCmd.Flags().IntP("port", "p", 3000, "port the server will bind")
	serverCmd.Flags().StringP("storage", "s", "memory", "storage type where data going to be persisted")
	serverCmd.Flags().Int("redis-port", 6379, "redis port for using with redis storage")
	serverCmd.Flags().IP("redis-host", net.IPv4(127, 0, 0, 1), "redis host for using with redis storage")

	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(displayCmd)
	rootCmd.AddCommand(serverCmd)
}
