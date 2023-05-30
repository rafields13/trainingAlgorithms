package main

import (
	"fmt"
	"sort"
)

type Shirt struct {
	Size  string
	Price float64
}

func biggerSize(a, b string) bool {
	if a == b {
		return false
	}

	numericalSize := map[string]int{
		"S":    1,
		"XS":   2,
		"M":    3,
		"XL":   4,
		"XXL":  5,
		"XXXL": 6,
	}

	_, aValid := numericalSize[a]
	_, bValid := numericalSize[b]
	if !aValid || !bValid {
		return false
	}

	return numericalSize[a] > numericalSize[b]
}

func calculateAverage(tshirts []Shirt) (float64, error) {
	if len(tshirts) == 0 {
		return 0, fmt.Errorf("error")
	}

	sort.SliceStable(tshirts, func(i, j int) bool {
		return biggerSize(tshirts[i].Size, tshirts[j].Size)
	})

	highestPrice := tshirts[0].Price
	lowestPrice := tshirts[len(tshirts)-1].Price

	average := (highestPrice + lowestPrice) / 2
	return average, nil
}

func main() {
	tshirts := []Shirt{
		{
			Size:  "M",
			Price: 10,
		},
		{
			Size:  "XXL",
			Price: 20,
		},
		{
			Size:  "S",
			Price: 7,
		},
		{
			Size:  "XXXXXXXS",
			Price: 5,
		},
	}

	average, err := calculateAverage(tshirts)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(average)
}
