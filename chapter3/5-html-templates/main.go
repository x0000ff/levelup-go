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
	Draft      bool
}

// Byline ...
func (a Article) Byline() string {
	return fmt.Sprintf("written by %s", a.AuthorName)
}

func main() {

	goArticle := Article{
		Name:       "The Go html/template package",
		AuthorName: "Mal Curtis",
		Draft:      false,
	}

	// tmpl, err := template.New("Foo").Parse("'{{ .Name }}'{{ if .Draft }} (Draft){{ end }} {{ .Byline }}")
	tmpl, err := template.New("Foo").Parse("'{{ .Name }}'{{ if .Draft }} (Draft){{ else }} (Published){{ end }} {{ .Byline }}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, goArticle)

	if err != nil {
		panic(err)
	}

}
