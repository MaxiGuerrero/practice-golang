package main

import "fmt"

func main(){
	literalInterpertateString := "Hola Mundo"
	NonLiteralInterpretateString := `Esta es una linea
	partida.`
	fmt.Println(literalInterpertateString)
	fmt.Printf("El tipo de literalString es: %T\n", literalInterpertateString)

	fmt.Println(NonLiteralInterpretateString)
	fmt.Printf("El tipo de literalString es: %T\n", NonLiteralInterpretateString)
}