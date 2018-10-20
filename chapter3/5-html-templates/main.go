package main

import (
	"html/template"
	"os"
)

func main() {

	tmpl, err := template.New("Foo").Parse("Price: ${{ . | printf \"%.2f\" }}\n")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, 12.3)

	if err != nil {
		panic(err)
	}

}
