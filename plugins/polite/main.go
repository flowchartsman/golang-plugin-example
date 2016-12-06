package main

import (
	"fmt"
	"github.com/alaska/golang-plugin-example/extern"
)

var PluginName = "polite"

var Action extern.ActionPlugin = func(p extern.Person) (string, error) {
	if p.Location == "back" || p.Location == "head" {
		return fmt.Sprintf("Gave %s a friendly pat on the %s", p.FirstName, p.Location), nil
	}
	return "", fmt.Errorf("Whoa there, buddy, that is NOT cool!")
}
