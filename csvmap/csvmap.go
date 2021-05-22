package csvmap

type CsvMap map[string]ItemCategory

type Product struct {
	SKU      string
	PLU      string
	Name     string
	Size     string
	SizeSort string
}

type ItemCategory struct {
	PLU   string
	name  string
	sizes []Item
}

type Item struct {
	SKU  string
	Size string
}

func (this CsvMap) AddItemSize(product Product, item Item) CsvMap {
	this[product.PLU] = this[product.PLU].add(item)

	return this
}

func (itemCat ItemCategory) add(item Item) ItemCategory {
	itemCat.sizes = append(itemCat.sizes, item)
	return itemCat
}

func (this CsvMap) CreateNewCategory(product Product, item Item) CsvMap {
	itemCategory := ItemCategory{
		PLU:   product.PLU,
		name:  product.Name,
		sizes: []Item{item},
	}

	this[product.PLU] = itemCategory

	return this
}

func (this CsvMap) IsInMap(product Product) bool {
	_, isInMap := this[product.PLU]
	return isInMap
}