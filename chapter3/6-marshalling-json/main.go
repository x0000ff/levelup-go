package main

import (
	"encoding/json"
	"fmt"
)

// Article ...
type Article struct {
	// Field appears in the JSON with the key "name"
	Name string `json:"name"`

	// Field appears in the JSON with the key "author_name"
	// but doesn't appear at all if its value is empty
	AuthorName string `json:"author_name,omitempty"`

	// Field will not appear in the JSON representation.
	ComissionPrice float64 `json:"-"`
}

func main() {
	article := Article{
		Name:           "JSON in Go",
		AuthorName:     "Mal Curtis",
		ComissionPrice: 23.1,
	}

	// data, err := json.Marshal(article)
	data, err := json.MarshalIndent(article, "", "   ")
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	} else {
		fmt.Println(string(data))
	}

}
