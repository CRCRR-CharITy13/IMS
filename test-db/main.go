package main

import (
	"fmt"
	data_utils "test-db/data/data-utils"
)

func main() {
	fmt.Printf("==== To Test The GIK-IMS Database\n")
	// itemFileName := "data/gik-ims-items.csv"
	locationFileName := "data/gik-ims-locations.csv"
	locations, err := data_utils.ReadCSV(locationFileName)
	if err != nil {
		fmt.Errorf("Error reading %s: %v\n", locationFileName, err)
	}

	fmt.Println("List of locations:\n")
	for _, location := range locations {
		fmt.Printf("%s : %s\n", location.Name, location.Description)
	}

}
