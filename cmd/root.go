package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

var (
	rootCmd = &cobra.Command{
		Use:   "ess",
		Short: "ESS(Ephemeral SSH Service) is a CLI tool to run an ephemeral SSH-able Fargate container easily",
		Long: `ESS is a short of "Ephemeral SSH Service".
You can run SSH-able container with this tool.
ESS works with ECR, ECS, Fargate and ssm-agent.
Mainly, Ephemeral SSH is useful for running temporary use server such as secure bastion.`,
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of ESS",
		Long:  `Print the version number of ESS`,
		Run:   cmdVersion,
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run a SSH-able container on Fargate",
		Long:  "Run a SSH-able container on Fargate",
		Run:   cmdRun,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(runCmd)
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}

func cmdVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("ess %s\n", version)
}

func cmdRun(cmd *cobra.Command, args []string) {
	fmt.Println("Run")
}
