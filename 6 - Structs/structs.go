package main

// Struct é um tipo de dado que permite armazenar diferentes tipos de dados em um único local.
type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	var usuarion1 = usuario{"Julia", 20, endereco{"Rua 1", 10}}
	println(usuarion1.nome)

	u2 := usuario{idade: 22}
	println(u2.idade)

	usuario3 := usuario{idade: 20}
	println(usuario3.nome)

}
