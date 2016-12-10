package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/marccampbell/kube-vault/config"
	"github.com/marccampbell/kube-vault/secret"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	yamlsEnv string
)

var yamlsCmd = &cobra.Command{
	Use:   "yamls",
	Short: "Generates secret yamls for all secrets in an environment",
	Long:  `Generates all of the secret yamls for an environment`,
}

func init() {
	viper.SetEnvPrefix("kubevault")
	viper.AutomaticEnv()

	yamlsCmd.PersistentFlags().StringVarP(&yamlsEnv, "env", "e", "", "The environment to generate secrets for.")

	yamlsCmd.RunE = yamls
}

// yamls is the implementation of the yamls command.
func yamls(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(yamlsCmd); err != nil {
		return err
	}

	//yamls := make([]string, 0, 0)
	env, err := config.GetEnvironment(yamlsEnv)
	if err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
		return nil
	}
	if env == nil {
		fmt.Printf("Environment %q not found\n", yamlsEnv)
		os.Exit(1)
		return nil
	}

	secrets, err := secret.ListSecretsInPath(*env.VaultAddress, *env.VaultToken, env.Path)
	if err != nil {
		fmt.Printf("Environment %q not found\n", yamlsEnv)
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
