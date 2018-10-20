package main

import (
	"html/template"
	"os"
)

func main() {

	tmpl, err := template.New("Foo").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "World")

	if err != nil {
		panic(err)
	}

}
