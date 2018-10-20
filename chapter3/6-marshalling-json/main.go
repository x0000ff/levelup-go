package main

import (
	"encoding/json"
	"fmt"
)

// Article ...
type Article struct {
	Name       string
	AuthorName string
	draft      bool
}

func main() {
	article := Article{
		Name:       "JSON in Go",
		AuthorName: "Mal Curtis",
		draft:      true,
	}

	// data, err := json.Marshal(article)
	data, err := json.MarshalIndent(article, "", "   ")
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	} else {
		fmt.Println(string(data))
	}

}
