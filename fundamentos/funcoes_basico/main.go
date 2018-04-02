package main

import ("fmt"
		"curso_de_go/fundamentos/funcoes_basico/matematica"
)

func main() {
	resultado := matematica.Somar(5, 5)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Subtrair(20, 5)
	fmt.Printf("O resultado da subtracao foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Multiplicar, 10, 10)
	fmt.Printf("O resultado da multiplicacao foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Dividir, 10, 5)
	fmt.Printf("O resultado da divisao foi: %d\r\n", resultado)
	resultado, resto := matematica.DividirComResto(12, 5) // Essa função retorna 2 valores
	fmt.Printf("O resultado da divisao foi: %d, e o resto é: %d\r\n", resultado, resto)
}

