package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "file.csv"
	comma := ";"
	delimiter(path, comma)
}

func delimiter(path string, delimiter string) (string, error) {

	csvFile, err := os.Open(path)
	if err != nil {
		return "", err
	}
	r := csv.NewReader(csvFile)
	r.Comma = ';'
	a := strconv.QuoteRune(r.Comma)

	fmt.Println(a)

	if strings.Contains(a, delimiter) {
		fmt.Println("La cadena contiene ';'")
	} else {
		fmt.Println("La cadena contiene ','")
	}
	return "", nil

}

// package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Open the CSV file
// 	f, err := os.Open("file.csv")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	// Create a new CSV reader
// 	r := csv.NewReader(f)

// 	// Set the delimiter to a comma
// 	r.Comma = ';'

// 	// Read in the CSV records
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Print out each record and field
// 	for _, record := range records {
// 		for _, field := range record {
// 			fmt.Printf("%s,", field)
// 		}
// 		fmt.Println()
// 	}
// }
