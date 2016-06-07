package main

import (
	"fmt"

	"github.com/PotentialWeb/iso3166/alpha2"
)

func main() {
	if alpha2.IsAlpha2("gb") {
		fmt.Println("gb is alpha2")
	}

	if !alpha2.IsAlpha2("foo") {
		fmt.Println("foo is not alpha2")
	}

	return
}