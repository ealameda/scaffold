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

var out io.Writer = os.Stdout

func main() {
	err := run()
	if err != nil {
		for _, e := range err {
			fmt.Fprintln(out, e.Error())
		}
	}
}

func run() []error {
	var allErrors []error

	var scaffold = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	scaffold.SetOutput(out)

	name := scaffold.String("n", "", "Project name")
	location := scaffold.String("l", "", "Project location on disk")
	url := scaffold.String("r", "", "Project remote repository url")
	static := scaffold.Bool("s", false, "Project will have static assets or not")
	err := scaffold.Parse(os.Args[1:])
	if err != nil {
		return nil
	}

	if *name == "" {
		allErrors = append(allErrors, errors.New("Project name cannot be empty"))
	}

	if *location == "" {
		allErrors = append(allErrors, errors.New("Project path cannot be empty"))
	}

	if *url == "" {
		allErrors = append(allErrors, errors.New("Project repository URL cannot be empty"))
	}
	if len(allErrors) > 0 {
		return allErrors
	}

	config := cfg{
		Name:     *name,
		Location: *location,
		Url:      *url,
		Static:   *static,
	}
	fmt.Fprintf(out, "Generating scaffold for project %s in %s\n", config.Name, config.Location)
	return nil

}
