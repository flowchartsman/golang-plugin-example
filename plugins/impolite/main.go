package main

import (
	"fmt"
	"github.com/alaska/golang-plugin-example/extern"
	"strings"
)

var PluginName = "impolite"

var Name extern.NamePlugin = func(p extern.Person) string {
	return fmt.Sprintf("More like Poop%s McJerk%s", strings.ToLower(p.FirstName), strings.ToLower(p.LastName))
}
