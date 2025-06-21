package main

import (
	"fmt"
)

func main() {
	var i int
	var n, n2 string

	fmt.Println("Olá, Estou aprendendo uma linguagem nova! ")
	fmt.Println("Olá, Estou aprendendo uma linguagem nova! ")
	fmt.Println("Olá, Estou aprendendo uma linguagem nova! ")
	fmt.Println("Olá, Estou aprendendo uma linguagem nova! ")

	fmt.Println("Digite seu nome: ")
	fmt.Scanln(&n, &n2)

	fmt.Println("Digite sua idade: ")
	fmt.Scan(&i)

	fmt.Println("Sua idade e seu nome são:", i, "e", n, n2)
}
