package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// Abrir el archivo CSV
	file, err := os.Open("file.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un nuevo lector CSV
	reader := csv.NewReader(file)

	// Leer la primera fila del archivo
	record, err := reader.Read()
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	// Imprimir la longitud de la fila y los caracteres que no sean letras o números
	fmt.Println("Longitud de la fila:", len(record))
	for _, char := range record[0] {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			fmt.Printf("Separador utilizado: %c\n", char)
		}
	}
}
