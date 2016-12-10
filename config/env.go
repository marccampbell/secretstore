package config

import (
	"errors"
	"strings"
)

// Environment represents a single environment we will
type Environment struct {
	Name string
	Path string

	VaultAddress *string
	VaultToken   *string
}

func (e *Environment) Create() error {
	if err := e.validate(); err != nil {
		return err
	}

	cfg, err := get()
	if err != nil {
		return err
	}

	cfg.Environments = append(cfg.Environments, e)

	if err = cfg.Save(); err != nil {
		return err
	}

	return nil
}

func (e *Environment) validate() error {
	// TODO make sure the environment doesn't already exist
	errs := make([]string, 0, 0)

	if e.Name == "" {
		errs = append(errs, "Name is required")
	}
	if e.Path == "" {
		errs = append(errs, "Path is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

// RemoveEnvironment will remove a cluster from the config file.
func RemoveEnvironment(name string) error {
	cfg, err := get()
	if err != nil {
		return err
	}

	envs := make([]*Environment, 0, 0)
	for _, env := range cfg.Environments {
		if env.Name != name {
			envs = append(envs, env)
		}
	}

	cfg.Environments = envs

	if err := cfg.Save(); err != nil {
		return err
	}

	return nil
}

func GetEnvironment(name string) (*Environment, error) {
	cfg, err := get()
	if err != nil {
		return nil, err
	}

	for _, env := range cfg.Environments {
		if env.Name == name {
			if env.VaultAddress == nil {
				env.VaultAddress = &cfg.VaultAddress
			}
			if env.VaultToken == nil {
				env.VaultToken = &cfg.VaultToken
			}

			return env, nil
		}
	}

	return nil, nil
}
