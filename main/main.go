package main

import (
	"encoding/csv"
	"fmt"
	"github.com/slatermorgan/csv-conv/csvmap"
	"github.com/slatermorgan/csv-conv/sizesorter"
	"io"
	"log"
	"os"
	"strings"
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
			Name:     removeQuotes(line[2]),
			Size:     removeQuotes(line[3]),
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

	fmt.Println(sizesorter.Sort(productMap).ToJSON())
}

func removeQuotes(str string) string {
	return strings.Replace(str, "\"", "", -1)
}
