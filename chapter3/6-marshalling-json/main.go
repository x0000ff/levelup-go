package main

import (
	"encoding/json"
	"fmt"
)

// Article ...
type Article struct {
	Name string `json:"name"`
}

// ArticleCollection ...
type ArticleCollection struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
}

func main() {
	article1 := Article{
		Name: "JSON in Go",
	}

	article2 := Article{
		Name: "Marshaling is easy",
	}

	articles := []Article{article1, article2}
	collection := ArticleCollection{
		Articles: articles,
		Total:    len(articles),
	}

	data, err := json.MarshalIndent(collection, "", "   ")
	if err != nil {
		fmt.Println("Couldn't marshal article:", err)
	} else {
		fmt.Println(string(data))
	}

}
