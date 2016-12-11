package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// KubeVaultCmd is the main (root) command for the CLI.
var KubeVaultCmd = &cobra.Command{
	Use:   "kube-vault",
	Short: "kube-vault helps manage Kubernetes Secrets in multiple environments by syncing from vault",
	Long:  "kube-vault manages Kubernetes Secrets in multiple environments. For more details, visit github.com/marccampbell/kube-vault",

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := InitializeConfig(); err != nil {
			return err
		}

		cmd.Usage()
		os.Exit(-1)
		return nil
	},
}

// Execute adds all child commands to the room command KubeVaultCmd and sets flags
func Execute() {
	AddCommands()

	if _, err := KubeVaultCmd.ExecuteC(); err != nil {
		os.Exit(-1)
	}
}

// AddCommands will add all child commands to the KubeVaultCmd
func AddCommands() {
	KubeVaultCmd.AddCommand(initCmd)
	KubeVaultCmd.AddCommand(envCmd)
	KubeVaultCmd.AddCommand(k8sCmd)
}

// InitializeConfig initializes the config environment with defaults.
func InitializeConfig(subCmdVs ...*cobra.Command) error {
	viper.SetEnvPrefix("kubevault")
	return nil
}
