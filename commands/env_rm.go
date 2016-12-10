package commands

import (
	"fmt"
	"os"

	"github.com/marccampbell/kube-vault/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	envRmName string
)

var envRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes a env",
	Long:  `Removes a env that is managed with kube-vault.`,
}

func init() {
	viper.SetEnvPrefix("kubevault")
	viper.AutomaticEnv()

	envRmCmd.PersistentFlags().StringVarP(&envRmName, "name", "n", "", "Name of the env")

	envRmCmd.RunE = envRm
}

// envRm is the implementation of the env rm command.
func envRm(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(envRmCmd); err != nil {
		return err
	}

	if err := config.RemoveEnvironment(envRmName); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
		return nil
	}

	return nil
}
