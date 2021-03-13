package main

import (
	"flag"
	"net"
	"strconv"

	"github.com/seadiaz/adoption/client/adoptables"
	"github.com/seadiaz/adoption/client/global"
	"github.com/seadiaz/adoption/client/memberships"
	"github.com/seadiaz/adoption/client/people"
	"github.com/seadiaz/adoption/client/teams"
	"github.com/seadiaz/adoption/client/utils"
	"github.com/seadiaz/adoption/server"
	"github.com/spf13/cobra"

	_ "github.com/seadiaz/adoption/docs"
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
	}

	loadPeopleCmd = &cobra.Command{
		Use:   "people",
		Short: "load people from a csv file",
		Long:  "load people from a csv file",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchLoadPeopleCommand,
	}

	loadTeamsCmd = &cobra.Command{
		Use:   "teams",
		Short: "load teams from a csv file",
		Long:  "load teams from a csv file",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchLoadTeamsCommand,
	}

	loadMembershipsCmd = &cobra.Command{
		Use:   "memberships <team>",
		Short: "load memberships from a csv file",
		Long:  "load memberships from a csv file",
		Args:  cobra.ExactArgs(1),
		Run:   dispatchLoadMembershipsCommand,
	}

	loadAdoptablesCmd = &cobra.Command{
		Use:   "adoptables",
		Short: "load adoptables from a csv file",
		Long:  "load adoptables from a csv file",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchLoadAdoptablesCommand,
	}

	displayCmd = &cobra.Command{
		Use:   "display",
		Short: "display the information stored in the server",
		Long:  "display the information stored in the server",
	}

	displayPeopleCmd = &cobra.Command{
		Use:   "people",
		Short: "display people information",
		Long:  "display people information",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchDisplayPeopleCommand,
	}

	displayTeamsCmd = &cobra.Command{
		Use:   "teams",
		Short: "display teams information",
		Long:  "display teams information",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchDisplayTeamsCommand,
	}

	displayMembershipsCmd = &cobra.Command{
		Use:   "memberships <tean>",
		Short: "display memberships information",
		Long:  "display memberships information",
		Args:  cobra.ExactArgs(1),
		Run:   dispatchDisplayMembershipsCommand,
	}

	displayAdoptablesCmd = &cobra.Command{
		Use:   "adoptables",
		Short: "display adoptables information",
		Long:  "display adoptables information",
		Args:  cobra.ExactArgs(0),
		Run:   dispatchDisplayAdoptablesCommand,
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

func dispatchLoadPeopleCommand(cmd *cobra.Command, args []string) {
	r := createPeopleRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Filename: cmd.Flag("file").Value.String(),
		Kind:     global.People,
		Action:   global.Load,
	}

	people.Execute(r, params)
}

func dispatchLoadAdoptablesCommand(cmd *cobra.Command, args []string) {
	r := createAdoptableRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Filename: cmd.Flag("file").Value.String(),
		Kind:     global.Adoptables,
		Action:   global.Load,
	}

	adoptables.Execute(r, params)
}

func createPeopleRepository(cmd *cobra.Command) *people.Repository {
	return people.CreateRepository(
		utils.CreateAPIClient(
			cmd.Flag("url").Value.String(),
			cmd.Flag("api-key").Value.String(),
		),
	)
}

func createAdoptableRepository(cmd *cobra.Command) *adoptables.Repository {
	return adoptables.CreateRepository(
		utils.CreateAPIClient(
			cmd.Flag("url").Value.String(),
			cmd.Flag("api-key").Value.String(),
		),
	)
}

func dispatchLoadTeamsCommand(cmd *cobra.Command, args []string) {
	r := createTeamsRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Filename: cmd.Flag("file").Value.String(),
		Kind:     global.Teams,
		Action:   global.Load,
	}

	teams.Execute(r, params)
}

func createTeamsRepository(cmd *cobra.Command) *teams.Repository {
	return teams.CreateRepository(
		utils.CreateAPIClient(
			cmd.Flag("url").Value.String(),
			cmd.Flag("api-key").Value.String(),
		),
	)
}

func dispatchLoadMembershipsCommand(cmd *cobra.Command, args []string) {
	r := createMembershipsRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Filename: cmd.Flag("file").Value.String(),
		Kind:     global.Memberships,
		Action:   global.Load,
		Parent:   args[0],
	}

	memberships.Execute(r, params)
}

func createMembershipsRepository(cmd *cobra.Command) *memberships.Repository {
	client := utils.CreateAPIClient(
		cmd.Flag("url").Value.String(),
		cmd.Flag("api-key").Value.String(),
	)
	return memberships.CreateRepository(
		people.CreateRepository(client),
		teams.CreateRepository(client),
	)
}

func dispatchDisplayPeopleCommand(cmd *cobra.Command, args []string) {
	repository := createPeopleRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Kind:   global.People,
		Action: global.Display,
	}

	people.Execute(repository, params)
}

func dispatchDisplayTeamsCommand(cmd *cobra.Command, args []string) {
	r := createTeamsRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Kind:   global.Teams,
		Action: global.Display,
	}

	teams.Execute(r, params)
}

func dispatchDisplayAdoptablesCommand(cmd *cobra.Command, args []string) {
	r := createAdoptableRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Kind:   global.Adoptables,
		Action: global.Display,
	}

	adoptables.Execute(r, params)
}

func dispatchDisplayMembershipsCommand(cmd *cobra.Command, args []string) {
	r := createMembershipsRepository(cmd)
	cmd.Flags()
	params := &global.CommandHandlerParams{
		Kind:   global.Memberships,
		Action: global.Display,
		Parent: args[0],
	}
	memberships.Execute(r, params)
}

func doBootServer(cmd *cobra.Command, args []string) {
	redisPort, _ := strconv.Atoi(cmd.Flag("redis-port").Value.String())
	params := &server.Params{
		Port:       cmd.Flag("port").Value.String(),
		Storage:    cmd.Flag("storage").Value.String(),
		RedisPort:  redisPort,
		RedisHost:  cmd.Flag("redis-host").Value.String(),
		BadgerPath: cmd.Flag("badger-path").Value.String(),
	}
	server.Boot(params)
}

func init() {
	loadCmd.PersistentFlags().StringP("url", "u", defaultURL, "The URL of the running instance of adoption server")
	loadCmd.PersistentFlags().StringP("api-key", "k", "", "API Key which is going to be send by Authorization header")
	loadCmd.PersistentFlags().StringP("file", "f", "", "Load data from `FILE` (required)")
	loadCmd.MarkFlagRequired("file")

	displayCmd.PersistentFlags().StringP("url", "u", defaultURL, "The URL of the running instance of adoption server")
	displayCmd.PersistentFlags().StringP("api-key", "k", "", "API Key which is going to be send by Authorization header")

	serverCmd.Flags().IntP("port", "p", 3000, "port the server will bind")
	serverCmd.Flags().StringP("storage", "s", "memory", "storage type where data going to be persisted")
	serverCmd.Flags().Int("redis-port", 6379, "redis port for using with redis storage")
	serverCmd.Flags().IP("redis-host", net.IPv4(127, 0, 0, 1), "redis host for using with redis storage")
	serverCmd.Flags().StringP("badger-path", "", "/tmp/badger", "badger path for using with badger storage")

	loadCmd.AddCommand(loadPeopleCmd)
	loadCmd.AddCommand(loadTeamsCmd)
	loadCmd.AddCommand(loadMembershipsCmd)
	loadCmd.AddCommand(loadAdoptablesCmd)
	displayCmd.AddCommand(displayPeopleCmd)
	displayCmd.AddCommand(displayTeamsCmd)
	displayCmd.AddCommand(displayMembershipsCmd)
	displayCmd.AddCommand(displayAdoptablesCmd)
	rootCmd.AddCommand(loadCmd)
	rootCmd.AddCommand(displayCmd)
	rootCmd.AddCommand(serverCmd)
}
