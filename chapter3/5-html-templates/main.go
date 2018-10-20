package main

import (
	"fmt"
	"html/template"
	"os"
)

// Article ...
type Article struct {
	Name       string
	AuthorName string
}

// Byline ...
func (a Article) Byline() string {
	return fmt.Sprintf("written by %s", a.AuthorName)
}

func main() {

	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
	}

	tmpl, err := template.New("Foo").Parse("'{{ .Name }}' {{ .Byline }}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, goArticle)

	if err != nil {
		panic(err)
	}

}
