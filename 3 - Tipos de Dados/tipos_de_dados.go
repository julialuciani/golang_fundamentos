package main

import (
	"errors"
	"fmt"
)

func main() {
	// os tipos de int no go são int8, int16, int32, int64
	// o tipo usado varia dependendo da necessidade de armazenamento
	// int8 = 8 bits
	// se for utilizado o int o próprio go vai utilizar o que faz mais sentido
	// os ints armazenam também valores negativos

	var numero int16 = 100
	fmt.Println(numero)

	//temos também o uint
	//significa unsigned int

	//var numero2 uint32 = -1000 //ele não deixa compilar esse valor

	//alias

	var numero3 rune = 124 //rune é um alias para int32
	fmt.Println(numero3)

	var numero2 byte = 123 //byte é um alias para uint8
	fmt.Println(numero2)

	//Números reais

	//float32, float64

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1233333333333.50
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45 //float64
	fmt.Println(numeroReal3)

	//Strings

	var str string = "Texto"
	fmt.Println(str)

	//Go não tem char, que seria um caracter só

	char := 'B'
	fmt.Println(char) //Se for um caracter só ele vai imprimir o valor da tabela ASCII, um número

	//Sempre utilizar aspas duplas pq ao usar aspas simples o go entende como Char

	//FIM Strings

	var texto int16
	fmt.Println(texto) //valor padrão de um int é 0 e de uma string é vazia

	var booleano1 bool = true //se deixar sem nada o valor é false
	fmt.Println(booleano1)

	//error

	var erro error
	fmt.Println(erro) //valor padrão é nil

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)

}
