package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type cfg struct {
	Name     string
	Location string
	Url      string
	Static   bool
}

func main() {
	config, err := setupScaffoldConfig(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	exitIfAnyErrors(validateConfig(config))

	generateScaffold(os.Stdout, config)
}

func validateConfig(config cfg) []error {
	var allErrors []error
	if config.Name == "" {
		allErrors = append(allErrors, errors.New("Project name cannot be empty"))
	}

	if config.Location == "" {
		allErrors = append(allErrors, errors.New("Project path cannot be empty"))
	}

	if config.Url == "" {
		allErrors = append(allErrors, errors.New("Project repository URL cannot be empty"))
	}

	return allErrors
}

func setupScaffoldConfig(out io.Writer, args []string) (cfg, error) {
	config := cfg{}
	scaffold := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	scaffold.SetOutput(out)
	scaffold.StringVar(&config.Name, "n", "", "Project name")
	scaffold.StringVar(&config.Location, "l", "", "Project location on disk")
	scaffold.StringVar(&config.Url, "r", "", "Project remote repository url")
	scaffold.BoolVar(&config.Static, "s", false, "Project will have static assets or not")
	err := scaffold.Parse(args)
	if err != nil {
		return config, err
	}
	if scaffold.NArg() != 0 {
		return config, errors.New("No positional parameters expected")
	}
	return config, err

}

func generateScaffold(out io.Writer, config cfg) {
	fmt.Fprintf(out, "Generating scaffold for project %s in %s\n", config.Name, config.Location)
}

func exitIfAnyErrors(errors []error) {
	if len(errors) != 0 {
		for _, e := range errors {
			fmt.Println(e)
		}
		os.Exit(1)
	}
}
