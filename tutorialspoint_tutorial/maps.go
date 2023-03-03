package main

import "fmt"

func main() {
	var country_capital_map map[string]string

	// This is how we make a map?
	country_capital_map = make(map[string]string)

	// Insert key-value pairs in the map
	country_capital_map["France"] = "Paris"
	country_capital_map["Italy"] = "Rome"
	country_capital_map["Japan"] = "Tokyo"
	country_capital_map["India"] = "New Delhi"

	// Print map using keys
	for country := range country_capital_map {
		fmt.Println("Capital of", country, "is", country_capital_map[country])
	}

	// Test if the entry is present in the map or not?
	capital, ok := country_capital_map["United States"]

	// If 'ok' is true, entry is present otherwise entry is absent.
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Capital of United States is not present.")
	}

	// Okay now we go and delete an entry
	delete(country_capital_map, "France")
	fmt.Println("Entry for Fance is deleted")

	fmt.Println("Updated map")

	// Print the map again.
	for country := range country_capital_map {
		fmt.Println("Capital of", country, "is", country_capital_map[country])
	}
}
