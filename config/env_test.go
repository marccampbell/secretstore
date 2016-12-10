package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateEnv(t *testing.T) {
	// verify a properly defined env passes
	env := Environment{
		Name: "name",
		Path: "path",
	}
	err := env.validate()
	require.NoError(t, err, "Correctly defined environment should pass validate")

	// verify that name is validated
	env = Environment{
		Name: "",
		Path: "path",
	}
	err = env.validate()
	require.Error(t, err, "")
	assert.Equal(t, err.Error(), "Name is required")

	// verify that path is validated
	env = Environment{
		Name: "name",
		Path: "",
	}
	err = env.validate()
	require.Error(t, err, "")
	assert.Equal(t, err.Error(), "Path is required")

	// verify that both are validated
	env = Environment{
		Name: "",
		Path: "",
	}
	err = env.validate()
	require.Error(t, err, "")
	assert.Equal(t, err.Error(), "Name is required, Path is required")
}
