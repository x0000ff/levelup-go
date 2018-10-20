package main

import (
	"html/template"
	"os"
)

// Multiply ...
func Multiply(a, b float64) float64 {
	return a * b
}

type Product struct {
	Price    float64
	Quantity float64
}

func main() {

	tmpl := template.New("Foo")
	tmpl.Funcs(template.FuncMap{
		"multiply": Multiply,
	})

	tmpl, err := tmpl.Parse("Price: ${{ multiply .Price .Quantity | printf \"%.2f\" }}\n")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, Product{
		Price:    12.3,
		Quantity: 3,
	})

	if err != nil {
		panic(err)
	}

}
