package main

import (
	"fmt"
	"github.com/alaska/golang-plugin-example/extern"
)

var PluginName = "boring"

var Name extern.NamePlugin = func(p extern.Person) string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

var Action extern.ActionPlugin = func(p extern.Person) (string, error) {
	return fmt.Sprintf("did nothing interesting with their %s", p.Location), nil
}
