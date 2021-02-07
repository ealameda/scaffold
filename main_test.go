package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validateConfig(t *testing.T) {
	t.Run("should return no errors with a valid config", func(t *testing.T) {
		config := cfg{Name: "test", Url: "http://www.test.com", Location: "./project"}
		errors := validateConfig(config)
		assert.Len(t, errors, 0)
	})

	t.Run("should return missing name if project name isn't set", func(t *testing.T) {
		config := cfg{Url: "http://www.test.com", Location: "./project"}
		errors := validateConfig(config)
		assert.Len(t, errors, 1)
		assert.Error(t, errors[0], "Project name cannot be empty")
	})

	t.Run("should return missing path if project location isn't set", func(t *testing.T) {
		config := cfg{Name: "test", Url: "http://www.test.com"}
		errors := validateConfig(config)
		assert.Len(t, errors, 1)
		assert.Error(t, errors[0], "Project path cannot be empty")
	})

	t.Run("should return missing location if project url isn't set", func(t *testing.T) {
		config := cfg{Name: "test", Location: "./project"}
		errors := validateConfig(config)
		assert.Len(t, errors, 1)
		assert.Error(t, errors[0], "Project repository URL cannot be empty")
	})

	t.Run("should return multiple errors with an empty config", func(t *testing.T) {
		config := cfg{}
		errors := validateConfig(config)
		assert.Len(t, errors, 3)
	})

}

func Test_generateScaffold(t *testing.T) {
	t.Run("should generate a message that the scaffold was generated based on the config", func(t *testing.T) {
		config := cfg{Name: "Project1", Url: "http://www.test.com", Location: "./project1"}
		byteBuf := new(bytes.Buffer)
		out := bufio.NewWriter(byteBuf)

		generateScaffold(out, config)
		out.Flush()
		assert.Equal(t, "Generating scaffold for project Project1 in ./project1\n", byteBuf.String())
	})
}

func Test_setupScaffoldConfig(t *testing.T) {
	t.Run("Should populate a config based on the input", func(t *testing.T) {
		args := []string{"-n", "project1", "-l", "/project1/", "-r", "github.com/ealameda/project1"}
		out := new(bytes.Buffer)
		//byteBuf.String()
		config, error := setupScaffoldConfig(out, args)
		assert.NoError(t, error)
		assert.Equal(t, "project1", config.Name)
		assert.Equal(t, "/project1/", config.Location)
		assert.Equal(t, "github.com/ealameda/project1", config.Url)
		assert.False(t, config.Static)
	})

	t.Run("should output help", func(t *testing.T) {
		args := []string{"-h"}
		out := new(bytes.Buffer)
		config, err := setupScaffoldConfig(out, args)
		assert.Error(t, err, "help requested")
		assert.Empty(t, config)
		assert.Contains(t, out.String(), "Usage of")

	})

	t.Run("should output errors", func(t *testing.T) {
		args := []string{"foo bar"}
		out := new(bytes.Buffer)
		config, err := setupScaffoldConfig(out, args)
		assert.Empty(t, config)
		assert.Error(t, err, "no positional arguments provided")

	})

}
