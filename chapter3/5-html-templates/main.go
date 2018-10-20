package main

import (
	"html/template"
	"os"
)

// Article ...
type Article struct {
	Name       string
	AuthorName string
}

func main() {

	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}

	tmpl, err := template.New("Foo").Parse("'{{ .Name }}' by {{ .AuthorName }}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, goArticle)

	if err != nil {
		panic(err)
	}

}
