package sizesorter

import (
	"fmt"
	"github.com/slatermorgan/csv-conv/csvmap"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Sort(csvmap csvmap.CsvMap) csvmap.CsvMap {

	for _, itemCategory := range csvmap {
		itemMap := itemCategory.Sizes
		sizeSort := itemCategory.SizeSort

		sort.SliceStable(itemMap, func(iIndex, jIndex int) bool {
			iSizeField := itemMap[iIndex].Size
			jSizeField := itemMap[jIndex].Size

			i, j := calcSortElements(iSizeField, jSizeField, sizeSort)

			return i < j
		})
	}

	return csvmap
}

func calcSortElements(
	iField string,
	jField string,
	sizeSort string,
) (float64, float64) {
	switch string(sizeSort) {
		case "SHOE_UK":
			return getUKSizeFloat(iField), getUKSizeFloat(jField)
		case "SHOE_EU":
			return getEUSizeFloat(iField), getEUSizeFloat(jField)
		case "CLOTHING_SHORT":
			return getClothingSizeFloat(iField), getClothingSizeFloat(jField)
		default:
			message := fmt.Sprintf(
				"Unsupported sort type: %v",
				sizeSort,
			)
			log.Fatal(message)

			return float64(0), float64(0)
	}
}

func sortShoeUK(itemMap []csvmap.Item) []csvmap.Item {

	sort.SliceStable(itemMap, func(i, j int) bool {
		return getUKSizeFloat(itemMap[i].Size) < getUKSizeFloat(itemMap[j].Size)
	})

	return itemMap
}
func getUKSizeFloat(size string) float64 {
	isChildSize := strings.Contains(size, "(Child)")

	if isChildSize {
		size = strings.Replace(size, " (Child)", "", -1)
	}

	float, err := strconv.ParseFloat(size, 64)

	if err != nil {
		log.Fatal(err)
	}

	// Apply 'child factor' used for sorting
	if isChildSize {
		return float / 100
	}

	return float
}
func getEUSizeFloat(size string) float64 {
	float, err := strconv.ParseFloat(size, 64)

	if err != nil {
		log.Fatal(err)
	}

	return float
}

func getClothingSizeFloat(size string) float64 {
	sizeLetterMap := map[string]float64{
		"XS":    1,
		"S":     2,
		"M":     3,
		"L":     4,
		"XL":    5,
		"XXL":   6,
		"XXXL":  7,
		"XXXXL": 8,
	}

	return sizeLetterMap[size]
}
