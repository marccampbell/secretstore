package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Manages environments",
	Long:  `kube-vault can be used to manage multiple environments`,
}

func init() {
	viper.SetEnvPrefix("kubevault")
	viper.AutomaticEnv()

	envCmd.AddCommand(envAddCmd)
	envCmd.AddCommand(envRmCmd)

	envCmd.RunE = env
}

// clusterI is the implementation of the cluster command.
func env(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(envCmd); err != nil {
		return err
	}

	cmd.Usage()

	return nil
}
