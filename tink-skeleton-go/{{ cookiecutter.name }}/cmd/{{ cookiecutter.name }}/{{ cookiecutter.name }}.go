// Having this file means that we will get a compiled binary with the name `{{ cookiecutter.name }}`.
package main

import "github.com/tink-ab/{{ cookiecutter.name }}/cmd"

func main() {
	cmd.Execute()
}
