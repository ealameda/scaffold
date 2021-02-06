package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestScaffold(t *testing.T) {
	out = bytes.NewBuffer(nil)

	t.Run("should display an information message if all fields are provided", func(t *testing.T) {
		os.Args = []string{"./scaffold", "-n", "Project1", "-l", "./project1", "-r", "github.com/username/project1"}
		out = bytes.NewBuffer(nil)
		main()
		assert.Equal(t, "Generating scaffold for project Project1 in ./project1\n", out.(*bytes.Buffer).String())
	})

	t.Run("should display errors if missing name", func(t *testing.T) {
		os.Args = []string{"./scaffold", "-l", "./project1", "-r", "github.com/username/project1"}
		out = bytes.NewBuffer(nil)
		main()
		assert.Equal(t, "Project name cannot be empty\n", out.(*bytes.Buffer).String())
	})

	t.Run("should display errors if missing location", func(t *testing.T) {
		os.Args = []string{"./scaffold", "-n", "Project1", "-r", "github.com/username/project1"}
		out = bytes.NewBuffer(nil)
		main()
		assert.Equal(t, "Project path cannot be empty\n", out.(*bytes.Buffer).String())
	})

	t.Run("should display errors if missing repository URL", func(t *testing.T) {
		os.Args = []string{"./scaffold", "-n", "Project1", "-l", "./project1"}
		out = bytes.NewBuffer(nil)
		main()
		assert.Equal(t, "Project repository URL cannot be empty\n", out.(*bytes.Buffer).String())
	})

	t.Run("should display multiple errors if both missing name and location", func(t *testing.T) {
		os.Args = []string{"./scaffold", "-r", "github.com/username/project1"}
		out = bytes.NewBuffer(nil)
		main()
		assert.Equal(t, "Project name cannot be empty\nProject path cannot be empty\n", out.(*bytes.Buffer).String())
	})

}
