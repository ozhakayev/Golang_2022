package main

import "fmt"

func main() {
	var count int = 22
	unitWeight := 0.4
	totalWeight := float64(count) * unitWeight
	fmt.Println(count, "cans weigh", totalWeight, "kilogram")
}