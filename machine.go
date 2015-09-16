package main

import (
	"github.com/flosch/pongo2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path"
	"strings"
)

// Machine configuration
type Machine struct {
	Hostname        string
	OperatingSystem string
	Finish          string
	Preseed         string
	ShortName       string
	Domain          string
	Network         []Interface
	Params          map[string]string
}

// Interface Configuration
type Interface struct {
	Name       string
	IPAddress  string
	MacAddress string
	Gateway    string
	Netmask    string
}

func machineDefinition(hostname string, machinePath string) (Machine, error) {
	var m Machine
	data, err := ioutil.ReadFile(path.Join(machinePath, hostname+".yaml"))
	if err != nil {
		return Machine{}, err
	}
	yaml.Unmarshal(data, &m)
	hostSlice := strings.Split(m.Hostname, ".")
	m.ShortName = hostSlice[0]
	m.Domain = strings.Join(hostSlice[1:], ".")
	return m, nil
}

func (m Machine) renderTemplate(templatePath string, config Config) (string, error) {
	var tpl = pongo2.Must(pongo2.FromFile(templatePath))
	result, err := tpl.Execute(pongo2.Context{"machine": m, "config": config})
	if err != nil {
		return "", err
	}
	return result, err
}