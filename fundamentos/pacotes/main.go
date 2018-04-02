package main

import ("fmt"
		"curso_de_go/fundamentos/pacotes/prefixo" // ou "./prefixo"
		"curso_de_go/fundamentos/pacotes/operadora") // ou "./operadora"

//NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Eduardo"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}