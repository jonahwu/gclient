package main

import (
	"fmt"
)

func IsApproach(dx float64, dx1 float64) bool {
	if dx <= dx1 {
		return true
	} else {
		return false
	}
	return true
}

func main() {
	dx := 6.0
	dx1 := 8.0
	if IsApproach(dx, dx1) && dx < 5 {
		fmt.Println("is apprached")
	}
}
