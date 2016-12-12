package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marccampbell/secretstore/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	initVaultAddr  string
	initVaultToken string
)
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the secretstore environment",
	Long:  `Initialize the secretstore environment by connecting to vault.`,
}

func init() {
	viper.SetEnvPrefix("secretstore")
	viper.AutomaticEnv()

	initCmd.PersistentFlags().StringVarP(&initVaultAddr, "vault-address", "a", "", "Vault address")
	initCmd.PersistentFlags().StringVarP(&initVaultToken, "vault-token", "t", "", "Vault access token")

	initCmd.RunE = initI
}

// initI is the implementation of the init command.
func initI(cmd *cobra.Command, args []string) error {
	if err := InitializeConfig(initCmd); err != nil {
		return err
	}

	// If keys already exist, we probably don't want to actually init, warn before proceeding
	if config.Exists() {
		fmt.Println("Config already exists. Cowardly refusing to overwrite it. Delete ~/.secretstore to re-init.")
		os.Exit(1)
		return nil
	}

	if len(initVaultAddr) == 0 || len(initVaultToken) == 0 {
		prompt := `
It's recommended to supply the Vault address and token using the --vault-address and --vault-token flags. 

Since you didn't provide these, the values will not be automaticalled inherited into each environment. You will have to supply these values when you create each environment. 
		
Are you sure you want to continue?`
		fmt.Printf("%s [y/n]: ", prompt)

		reader := bufio.NewReader(os.Stdin)
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%#v", err)
			os.Exit(1)
			return nil
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response != "y" && response != "yes" {
			os.Exit(1)
			return nil
		}
	}

	cfg := config.Config{
		SchemaVersion: 1,
		VaultAddress:  initVaultAddr,
		VaultToken:    initVaultToken,
	}

	if err := cfg.Save(); err != nil {
		fmt.Printf("%#v\n", err)
		os.Exit(1)
	}

	return nil
}
