package main

import (
	"fmt"
	"os"
	"text/template"

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
	networks := []Network{
		{
			Name:           "test1",
			Gw_v4:          "172.21.77.1",
			Network_v4:     "172.21.77.0/24",
			Interface_name: "eth1",
		},
		{
			Name:           "test2",
			Gw_v4:          "172.21.21.1",
			Network_v4:     "172.21.21.0/24",
			Interface_name: "eth2",
		},
	}
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return errors.Wrap(err, "parsing template file")
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return errors.Wrap(err, "creating output file")
	}
	defer output.Close()

	err = tmpl.Execute(output, networks)
	if err != nil {
		return errors.Wrap(err, "executing template file")
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
