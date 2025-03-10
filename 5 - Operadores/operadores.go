package main

import "fmt"

func main() {
	//Aritméticos

	soma := 1 + 2
	subtracao := 1 - 2
	multiplicacao := 1 * 2
	divisao := 1 / 2
	restoDaDivisao := 1 % 2

	println(soma)
	println(subtracao)
	println(multiplicacao)
	println(divisao)
	println(restoDaDivisao)

	var numero1 int16 = 10

	var numero2 int16 = 25
	soma2 := numero1 + int16(numero2)

	println(soma2)

	//Atribuição

	var variavel1 string = "String"
	variavel2 := "String2"

	println(variavel1)
	println(variavel2)

	//Relacionais

	println(1 > 2)
	println(1 >= 2)
	println(1 == 2)
	println(1 <= 2)
	println(1 < 2)
	println(1 != 2)

	//Lógicos
	vardadeiro, falso := true, false
	fmt.Println(vardadeiro && falso)
	fmt.Println(vardadeiro || falso)
	fmt.Println(!vardadeiro)

	//Unários
	numero := 10
	numero++
	numero += 15
	println(numero)

	numero--
	numero -= 20
	println(numero)

	//Ternário
	//Não existe em Go

	var texto string

	if numero == 5 {
		texto = "Igual 5"
	} else if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	println(texto)
	println("O número é ", numero)

}
