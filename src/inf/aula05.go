package main

import "fmt"


func main() {
	var nome string = "Matias"
	var idade int = 29
	var altura float64 = 1.75
	var versao float32 = 1.1

	fmt.Println("Olá meu Nome é: ", nome, " e tenho ", idade, " anos e ", altura, " de altura")
	fmt.Println("Este programa está na versão ", versao)
}