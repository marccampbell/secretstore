package commands

import (
	"fmt"
	"os"

	"github.com/marccampbell/secretstore/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	envAddName string
	envAddPath string

	envAddVaultAddr  string
	envAddVaultToken string
)

var envAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new environment",
	Long:  `Adds a new environment that can be managed with secretstore.`,
}

func init() {
	viper.SetEnvPrefix("secretstore")
	viper.AutomaticEnv()

	envAddCmd.PersistentFlags().StringVarP(&envAddName, "name", "n", "", "Name of the env")
	envAddCmd.PersistentFlags().StringVarP(&envAddPath, "path", "p", "", "The path in vault where the secrets are stored")
	envAddCmd.PersistentFlags().StringVarP(&envAddVaultAddr, "vault-address", "a", "", "(optional) The overridden vault address to get the secrets for this environment")
	envAddCmd.PersistentFlags().StringVarP(&envAddVaultToken, "vault-token", "t", "", "(optional) The overridden token to use when accessing the vault")

	envAddCmd.RunE = envAdd
}

// envAdd is the implementation of the env add command.
func envAdd(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(envAddCmd); err != nil {
		return err
	}

	env := config.Environment{
		Name: envAddName,
		Path: envAddPath,
	}

	if len(envAddVaultAddr) > 0 {
		env.VaultAddress = &envAddVaultAddr
	}
	if len(envAddVaultToken) > 0 {
		env.VaultToken = &envAddVaultToken
	}

	if err := env.Create(); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
		return nil
	}

	return nil
}
