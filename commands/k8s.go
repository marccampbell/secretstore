package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/marccampbell/secretstore/config"
	"github.com/marccampbell/secretstore/secret"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	k8sEnv string
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Generates kubernetes yaml for all secrets in an environment",
	Long:  `Generates all of the kubernetes yaml for an environment`,
}

func init() {
	viper.SetEnvPrefix("secretstore")
	viper.AutomaticEnv()

	k8sCmd.PersistentFlags().StringVarP(&k8sEnv, "env", "e", "", "The environment to generate secrets for.")

	k8sCmd.RunE = k8s
}

// k8s is the implementation of the k8s command.
func k8s(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(k8sCmd); err != nil {
		return err
	}

	env, err := config.GetEnvironment(k8sEnv)
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
		return nil
	}
	if env == nil {
		fmt.Printf("Environment %q not found\n", k8sEnv)
		os.Exit(1)
		return nil
	}

	secrets, err := secret.ListSecretsInPath(*env.VaultAddress, *env.VaultToken, env.Path)
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
		return nil
	}

	results := make([]string, 0, 0)
	for _, s := range secrets {
		results = append(results, s.CreateYAML())
	}

	fmt.Println(strings.Join(results, "---\n"))

	return nil
}
