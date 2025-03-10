package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"
	fmt.Println(variavel2)
	fmt.Println(variavel1)

	var (
		variavel3 string = "ju"
		variavel4 string = "gui"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"

	fmt.Println(variavel5, variavel6)

	const contante1 string = "constante 1"
	fmt.Println(contante1)

	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel5, variavel6)

}
