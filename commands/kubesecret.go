package commands

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SecretStoreCmd is the main (root) command for the CLI.
var SecretStoreCmd = &cobra.Command{
	Use:   "secretstore",
	Short: "secretstore helps manage Kubernetes Secrets in multiple environments by syncing from vault",
	Long:  "secretstore manages Kubernetes Secrets in multiple environments. For more details, visit github.com/marccampbell/secretstore",

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := InitializeConfig(); err != nil {
			return err
		}

		cmd.Usage()
		os.Exit(-1)
		return nil
	},
}

// Execute adds all child commands to the room command SecretStoreCmd and sets flags
func Execute() {
	AddCommands()

	if _, err := SecretStoreCmd.ExecuteC(); err != nil {
		os.Exit(-1)
	}
}

// AddCommands will add all child commands to the SecretStoreCmd
func AddCommands() {
	SecretStoreCmd.AddCommand(initCmd)
	SecretStoreCmd.AddCommand(envCmd)
	SecretStoreCmd.AddCommand(k8sCmd)
}

// InitializeConfig initializes the config environment with defaults.
func InitializeConfig(subCmdVs ...*cobra.Command) error {
	viper.SetEnvPrefix("secretstore")
	return nil
}
