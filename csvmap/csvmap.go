package csvmap

import (
	"encoding/json"
    "fmt"
)

type CsvMap map[string]ItemCategory

type Product struct {
	SKU      string
	PLU      string
	Name     string
	Size     string
	SizeSort string
}

type ItemCategory struct {
	PLU   string `json: "PLU"`
	Name  string `json: "name"`
	Sizes []Item `json: "sizes"`
}

type Item struct {
	SKU  string `json: "SKU"`
	Size string `json: "size"`
}

func (this CsvMap) AddItemSize(product Product, item Item) CsvMap {
	this[product.PLU] = this[product.PLU].add(item)

	return this
}

func (itemCat ItemCategory) add(item Item) ItemCategory {
	itemCat.Sizes = append(itemCat.Sizes, item)
	return itemCat
}

func (this CsvMap) CreateNewCategory(product Product, item Item) CsvMap {
	itemCategory := ItemCategory{
		PLU:   product.PLU,
		Name:  product.Name,
		Sizes: []Item{item},
	}

	this[product.PLU] = itemCategory

	return this
}

func (this CsvMap) IsInMap(product Product) bool {
	_, isInMap := this[product.PLU]
	return isInMap
}

func (this CsvMap) ToJSON() string {
	json, err := json.Marshal(this)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	return string(json)
}