package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type Shirt struct {
	Size  string
	Price float64
}

func average(shirts []Shirt) float64 {
	var sum float64

	for _, shirt := range shirts {
		sum += shirt.Price
	}

	return sum / float64(len(shirts))
}
func CalculateAveragePrice(shirts []Shirt) (max, min float64, err error) {
	if len(shirts) == 0 {
		return 0, 0, errors.New("erro")
	}

	shirtsMap := make(map[int][]Shirt)

	var maxSize = math.MinInt
	var minSize = math.MaxInt

	for _, shirt := range shirts {
		var result int
		switch {
		case strings.Contains(shirt.Size, "S"):
			result = -strings.Count(shirt.Size, "X") - 1
		case strings.Contains(shirt.Size, "L"):
			result = strings.Count(shirt.Size, "X") + 1

		}
		shirtsMap[result] = append(shirtsMap[result], shirt)
		if result > maxSize {
			maxSize = result
		}
		if result < minSize {
			minSize = result
		}
	}

	return average(shirtsMap[maxSize]), average(shirtsMap[minSize]), nil
}

func main() {
	shirts := []Shirt{
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

	min, max, err := CalculateAveragePrice(shirts)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(min, max)
	}
}
