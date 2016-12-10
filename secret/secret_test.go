package secret

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateYAMLSingleSecret(t *testing.T) {
	oneSecret := make(map[string]interface{})
	oneSecret["username"] = "admin"

	s := Secret{
		Name:    "one",
		Secrets: oneSecret,
	}

	valueEncoded := base64.StdEncoding.EncodeToString([]byte("admin"))
	correctYAML := `apiVersion: v1
kind: Secret
metadata:
  name: one
type: Opaque
data:
  username: ` + valueEncoded + "\n"
	assert.Equal(t, correctYAML, s.CreateYAML())

}

func TestCreateYAMLMultipleSecrets(t *testing.T) {
	twoSecrets := make(map[string]interface{})
	twoSecrets["username"] = "admin"
	twoSecrets["password"] = "topsecret"

	s := Secret{
		Name:    "two",
		Secrets: twoSecrets,
	}

	valueEncodedFirst := base64.StdEncoding.EncodeToString([]byte("admin"))
	valueEncodedSecond := base64.StdEncoding.EncodeToString([]byte("topsecret"))
	correctYAML := `apiVersion: v1
kind: Secret
metadata:
  name: two
type: Opaque
data:
  username: ` + valueEncodedFirst + `
  password: ` + valueEncodedSecond + "\n"

	assert.Equal(t, correctYAML, s.CreateYAML())
}
