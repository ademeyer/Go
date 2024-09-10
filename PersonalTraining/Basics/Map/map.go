package main

import (
	"fmt"
	"sort"
)

func main() {
	countries_initials := []string{"NG", "GY", "UK", "US", "EG", "GH"}
	countries_names := []string{"Nigeria", "Germany", "United Kingdom", "United Stated", "Egypt", "Ghana"}
	countries := make(map[string]string) // intial, name

	for i, val := range countries_names {
		countries[countries_initials[i]] = val
	}

	for key, val := range countries {
		fmt.Printf("%v: %v\n", key, val)
	}

	// Try and print map in a sorted way
	fmt.Println("\nSorting maps: ")
	Keys := make([]string, len(countries))
	i := 0
	for k := range countries {
		Keys[i] = k
		i++
	}
	sort.Strings(Keys)
	for i := range Keys {
		fmt.Printf("%v: %v\n", Keys[i], countries[Keys[i]])
	}

	items := []string{"apple", "banana", "orange", "date"}
	fmt.Println((convertToMap(items)))
}

// Converts a slice of strings to a map object.
func convertToMap(items []string) map[string]float64 {

	num := len(items)

	// Create a map object
	result := make(map[string]float64)
	// Your code goes here
	for i := 0; i < num; i++ {
		result[items[i]] = (100 / float64(num))
	}
	return result
}
