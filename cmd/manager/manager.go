package manager

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/camptocamp/bivac/cmd"
	"github.com/camptocamp/bivac/internal/manager"
	"github.com/camptocamp/bivac/pkg/volume"
)

var (
	// Server stores informations relative the Bivac server
	server manager.Server

	// Orchestrator is the name of the orchestrator on which Bivac should connect to
	Orchestrator string

	Orchestrators manager.Orchestrators

	dbPath string
)
var envs = make(map[string]string)

// TODO: Rename this command to something more explicit
var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Start Bivac backup manager",
	Run: func(cmd *cobra.Command, args []string) {
		// Global variables
		whitelistVolumes, _ := cmd.Flags().GetString("whitelist")
		//blacklistVolumes, _ := cmd.Flags().GetString("blacklist")

		volumesFilters := volume.Filters{
			Whitelist: strings.Split(whitelistVolumes, ","),
		}

		o, err := manager.GetOrchestrator(Orchestrator, Orchestrators)
		if err != nil {
			log.Errorf("failed to retrieve orchestrator: %s", err)
			return
		}

		err = manager.Start(o, server, volumesFilters)
		if err != nil {
			log.Errorf("failed to start manager: %s", err)
			return
		}
	},
}

func init() {
	managerCmd.Flags().StringVarP(&server.Address, "server.address", "", "0.0.0.0:8182", "Address to bind on.")
	envs["BIVAC_SERVER_ADDRESS"] = "server.address"
	managerCmd.Flags().StringVarP(&server.PSK, "server.psk", "", "", "Pre-shared key.")
	envs["BIVAC_SERVER_PSK"] = "server.psk"

	managerCmd.Flags().StringVarP(&Orchestrator, "orchestrator", "o", "", "Orchestrator on which Bivac should connect to.")
	envs["BIVAC_ORCHESTRATOR"] = "orchestrator"

	managerCmd.Flags().StringVarP(&Orchestrators.Docker.Endpoint, "docker.endpoint", "", "unix:///var/run/docker.sock", "Docker endpoint.")
	envs["BIVAC_DOCKER_ENDPOINT"] = "docker.endpoint"

	managerCmd.Flags().StringVarP(&dbPath, "db.path", "", "bivac.sqlite", "Path the Bivac internal database.")
	envs["BIVAC_DB_PATH"] = "db.path"

	cmd.SetValuesFromEnv(envs, managerCmd.Flags())
	cmd.RootCmd.AddCommand(managerCmd)
}
