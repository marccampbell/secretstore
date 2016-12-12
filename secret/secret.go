package secret

import (
	"encoding/base64"
	"fmt"
	"path"
	"strings"

	"github.com/marccampbell/secretstore/log"

	vault "github.com/hashicorp/vault/api"
)

type Secret struct {
	Name    string
	Secrets map[string]interface{}
}

// CreateYAML will create a new secret for the given clusterName
func (s *Secret) CreateYAML() string {
	vals := make([]string, 0, 0)
	for name, value := range s.Secrets {
		encoded := base64.StdEncoding.EncodeToString([]byte(value.(string)))
		vals = append(vals, fmt.Sprintf("  %s: %s", name, encoded))
	}

	return fmt.Sprintf(newSecretTemplate, s.Name, strings.Join(vals, "\n"))
}

func ListSecretsInPath(address string, token string, vaultPath string) ([]*Secret, error) {
	vaultConfig := vault.DefaultConfig()
	vaultConfig.Address = address

	vaultClient, err := vault.NewClient(vaultConfig)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	vaultClient.SetToken(token)

	vaultSecrets, err := vaultClient.Logical().List(vaultPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	secrets := make([]*Secret, 0, 0)

	secretNames := vaultSecrets.Data["keys"].([]interface{})
	for _, secretName := range secretNames {
		items, err := getSecretsInItem(vaultClient, path.Join(vaultPath, secretName.(string)))
		if err != nil {
			log.Error(err)
			return nil, err
		}

		secret := Secret{
			Name:    secretName.(string),
			Secrets: items,
		}
		secrets = append(secrets, &secret)
	}

	return secrets, nil
}

func getSecretsInItem(vaultClient *vault.Client, vaultPath string) (map[string]interface{}, error) {
	item, err := vaultClient.Logical().Read(vaultPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return item.Data, nil
}
