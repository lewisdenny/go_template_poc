package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/go-sprout/sprout"
	"github.com/go-sprout/sprout/registry/std"
	"github.com/pkg/errors"
)

type Network struct {
	Name           string `yaml:"name"`
	Gw_v4          string `yaml:"gw_v4"`
	Network_v4     string `yaml:"network_v4"`
	Interface_name string `yaml:"interface_name"`
}

func run() error {
	const templateFile = "template/template.tmpl"
	const outputFile = "parsed.yaml"

	// Data to render into template, could be read from file
	networks := []Network{
		{
			Name:       "test1",
			Gw_v4:      "172.21.77.1",
			Network_v4: "172.21.77.0/24",
			// Interface_name: "eth1",
		},
		{
			Name:       "test2",
			Gw_v4:      "172.21.21.1",
			Network_v4: "172.21.21.0/24",
			// Interface_name: "eth2",
		},
	}

	// Create Spront handler
	handler := sprout.New()
	handler.AddRegistry(std.NewRegistry())
	funcs := handler.Build()

	f, _ := os.ReadFile(templateFile)

	// Create template passing in spront functions
	tmpl, err := template.New("template.tmpl").Funcs(funcs).Parse(string(f))
	if err != nil {
		return errors.Wrap(err, "parsing template file")
	}

	// Create output file
	output, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrap(err, "creating output file")
	}
	defer output.Close()

	// Debug
	err = tmpl.Execute(os.Stdout, networks)
	if err != nil {
		return errors.Wrap(err, "executing template file")
	}

	// Render template
	// err = tmpl.Execute(output, networks)
	// if err != nil {
	// 	return errors.Wrap(err, "executing template file")
	// }

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
