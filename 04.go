package main

import (
	"fmt"
	"strings"
)

type Tshirt struct {
	Price int
	Size  string
}

func averageHighestLowestPrice(tshirts []Tshirt) int {
	var bigger Tshirt
	var smaller Tshirt

	for _, tshirt := range tshirts {
		if strings.HasSuffix(tshirt.Size, "L") && strings.Count(tshirt.Size, "")-1 > strings.Count(bigger.Size, "")-1 {
			tshirt.Price = bigger.Price
		}
	}
	for _, tshirt := range tshirts {
		if strings.HasSuffix(tshirt.Size, "S") && strings.Count(tshirt.Size, "")-1 > strings.Count(bigger.Size, "")-1 {
			tshirt.Price = smaller.Price
		}
	}

	return (bigger.Price + smaller.Price) / 2
}

func main() {
	tshirts := []Tshirt{
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

	calculating := averageHighestLowestPrice(tshirts)
	fmt.Print(calculating)
}
