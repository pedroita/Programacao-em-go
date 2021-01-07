package main

import (
	"fmt"
	"os"
)

func main() {
	var unidadeDestino string
	var unidadeOringem string
	if len(os.Args) < 3 {
		fmt.Println("Uso:converso<valores><unidade>")
		os.Exit(1)
	}
	if unidadeOringem == "celsius" {
		unidadeDestino = "fahrenheit"
	} else if unidadeOringem == "quilometros" {
		unidadeDestino = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida!", unidadeDestino)
		os.Exit(1)
	}
}
