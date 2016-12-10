package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/user"
	"path"

	"github.com/marccampbell/kube-vault/log"

	"github.com/BurntSushi/toml"
)

// Config represents the format of the config file.
type Config struct {
	SchemaVersion int

	VaultAddress string
	VaultToken   string

	Environments []*Environment
}

// Exists returns a bool if a current config exists
func Exists() bool {
	usr, err := user.Current()
	if err != nil {
		log.Error(err)
		return false
	}

	configPath := path.Join(usr.HomeDir, ".kube-vault", "config.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	}

	return true
}

// Save will serialize the config file to disk.
func (c *Config) Save() error {
	usr, err := user.Current()
	if err != nil {
		log.Error(err)
		return err
	}

	configPath := path.Join(usr.HomeDir, ".kube-vault")
	if err := os.MkdirAll(configPath, 0755); err != nil {
		log.Error(err)
		return err
	}

	configPath = path.Join(configPath, "config.toml")

	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		log.Error(err)
		return err
	}

	if err := ioutil.WriteFile(configPath, buf.Bytes(), 0644); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func get() (*Config, error) {
	usr, err := user.Current()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	configPath := path.Join(usr.HomeDir, ".kube-vault", "config.toml")

	cfgData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	cfg := Config{}
	if _, err := toml.Decode(string(cfgData), &cfg); err != nil {
		log.Error(err)
		return nil, err
	}

	return &cfg, nil
}
