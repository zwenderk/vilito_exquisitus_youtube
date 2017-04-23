package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Introducción a los arreglos")
	_="breakpoint"
	printTestResults()
}

func printTestResults() {
	var testResult [10]int
	testResult[1] = 12
	fmt.Println("Resultado de prueba 1", testResult)
}
