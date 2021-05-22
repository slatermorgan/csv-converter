package main

import (
    "fmt"
	"encoding/csv"
	"io"
	"log"
	"os"
    "github.com/slatermorgan/csv-conv/csvmap"
)

func main() {
	file, err := os.Open("products.csv")
	reader := csv.NewReader(file)
	reader.LazyQuotes = true

	if err != nil {
		fmt.Println(err)
	}

	// Associative array [ PLU => ItemCategory ]
	productMap := make(csvmap.CsvMap)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		product := csvmap.Product{
			SKU:      line[0],
			PLU:      line[1],
			Name:     line[2],
			Size:     line[3],
			SizeSort: line[4],
		}

		item := csvmap.Item{
			SKU:  product.SKU,
			Size: product.Size,
		}

		if productMap.IsInMap(product) {
			productMap.AddItemSize(product, item)
		} else {
			productMap.CreateNewCategory(product, item)
		}
	}

	fmt.Println(productMap)
	// fmt.Println(csvmap.toJSON())
}