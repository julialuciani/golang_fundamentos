package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	auxiliar.Escrever()
	auxiliar.Escrever2()
	fmt.Println("Escrevendo no main")

	erro := checkmail.ValidateFormat("julialucianiu.u@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("kskskwk")
	fmt.Println(erro2)
}
