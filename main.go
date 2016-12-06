package main

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/alaska/golang-plugin-example/extern"
)

var NamePlugins = map[string]extern.NamePlugin{}
var ActionPlugins = map[string]extern.ActionPlugin{}

func loadPlugin(path string, f os.FileInfo, _ error) error {
	if f.IsDir() {
		return nil
	}
	p, err := plugin.Open(path)
	if err != nil {
		fmt.Printf("invalid plugin: %s", path)
		return nil
	}
	var pluginName string
	if ns, err := p.Lookup("PluginName"); err == nil {
		if n, ok := ns.(*string); ok {
			pluginName = *n
		}
	}

	if pluginName == "" {
		fmt.Printf("plugin %s does not register a name\n", path)
		return nil
	}

	fmt.Printf("%s: plugin [%s]\n", path, pluginName)

	if s, err := p.Lookup("Name"); err == nil {
		if npf, ok := s.(*extern.NamePlugin); ok {
			fmt.Println("\tprovides name plugin")
			NamePlugins[pluginName] = *npf
		} else {
			fmt.Println("\tdoes not correctly implement name plugin")
		}
	}

	if s, err := p.Lookup("Action"); err == nil {
		if apf, ok := s.(*extern.ActionPlugin); ok {
			fmt.Println("\tprovides action plugin")
			ActionPlugins[pluginName] = *apf
		} else {
			fmt.Println("\tdoes not correctly implement action plugin")
		}
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please provide a directory with plugins as the first argument")
		os.Exit(1)
	}

	pluginDir := os.Args[1]

	if fi, err := os.Stat(pluginDir); err != nil || !fi.IsDir() {
		fmt.Println("plugin directory does not exist or is not directory")
		os.Exit(1)
	}

	err := filepath.Walk(pluginDir, loadPlugin)
	if err != nil {
		panic(err)
	}
	fmt.Printf("----------\n")
	people := []extern.Person{
		{
			FirstName: "Andy",
			LastName:  "Walker",
			Location:  "head",
		},
		{
			FirstName: "Bob",
			LastName:  "Everyman",
			Location:  "knee",
		},
	}

	for i, p := range people {
		fmt.Printf("person %d\n", i)
		fmt.Println("\tname plugins")
		for k, v := range NamePlugins {
			fmt.Printf("\t\t[%s]: %s\n", k, v(p))
		}
		fmt.Println("\taction plugins")
		for k, v := range ActionPlugins {
			result, err := v(p)
			if err != nil {
				fmt.Printf("\t\t[%s]: ERROR: %s\n", k, err.Error())
			} else {
				fmt.Printf("\t\t[%s]: %s\n", k, result)
			}
		}
	}
}
