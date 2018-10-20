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

	tmpl, err := template.New("Foo").Parse(`
    {{ define "ArticleResource" }}
      <p>'{{ .Name }}' by {{ .AuthorName }}</p>
    {{ end }}

    {{ define "ArticleLoop" }}
      {{ range . }}
        {{ template "ArticleResource" . }}
      {{ else }}
        <p>No published articles yet</p>
      {{ end }}
    {{ end }}

    {{ template "ArticleLoop" . }}
    `)

	articles := []Article{
		Article{
			Name:       "The Go html/template package",
			AuthorName: "Mal Curtis",
			Draft:      false,
		},
		Article{
			Name:       "Code Complete",
			AuthorName: "Steve McConnell",
			Draft:      false,
		},
	}
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, articles)

	if err != nil {
		panic(err)
	}

}
