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
				sortClothingSort(itemCategory.Sizes)
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
		iSizeInt, err := strconv.Atoi(strings.TrimSpace(itemMap[i].Size))
		jSizeInt, err := strconv.Atoi(strings.TrimSpace(itemMap[j].Size))

		if err != nil {
			fmt.Println(err)
		}

    	return iSizeInt < jSizeInt
	})

	return itemMap
}
func sortClothingSort(itemMap []csvmap.Item) []csvmap.Item {
	sizeLetterMap := map[string]int{
		"XS": 1,
		"S": 2,
		"M": 3,
		"L": 4,
		"XL": 5,
		"XXL": 6,
		"XXXL": 7,
		"XXXXL": 8,
	}

	sort.SliceStable(itemMap, func(i, j int) bool {
		iSize := strings.TrimSpace(itemMap[i].Size)
		jSize := strings.TrimSpace(itemMap[j].Size)

    	return sizeLetterMap[iSize] < sizeLetterMap[jSize]
	})

	return itemMap
}