package sizesorter

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
	"github.com/slatermorgan/csv-conv/csvmap"
)

// Hello returns a greeting for the named person.
func Sort(csvmap csvmap.CsvMap) csvmap.CsvMap {
	for _, itemCategory := range csvmap {

		switch string(itemCategory.SizeSort) {
			case " SHOE_UK":
				sortShoeUK(itemCategory.Sizes)
			case " SHOE_EU":
				sortShoeEU(itemCategory.Sizes)
			case " CLOTHING_SHORT":
				sortClothingShort(itemCategory.Sizes)
			default:
				message := fmt.Sprintf(
					"Unsupported sort type: %v",
					itemCategory.SizeSort,
				)
    			fmt.Println(message)
		}
    }

    return csvmap
}

func sortShoeUK(itemMap []csvmap.Item) []csvmap.Item {

    fmt.Println("sortShoeUK")
	return itemMap
}
func sortShoeEU(itemMap []csvmap.Item) []csvmap.Item {

	sort.SliceStable(itemMap, func(i, j int) bool {
		iSize, err := strconv.Atoi(strings.TrimSpace(itemMap[i].Size))
		jSize, err := strconv.Atoi(strings.TrimSpace(itemMap[j].Size))

		if err != nil {
			fmt.Println(err)
		}

    	return iSize < jSize
	})

	return itemMap
}
func sortClothingShort(itemMap []csvmap.Item) []csvmap.Item {

    fmt.Println("sortClothingShort")
	return itemMap
}