package matematica

func Multiplicar(x int, y int) int {
	return x * y
}

func Dividir(x int, y int) int {
	return x / y
}

func Somar(x int, y int) (resultado int) {
	resultado = x + y
	return
}

func Subtrair(x int, y int) (resultado int) {
	resultado = x - y
	return 
}

func DividirComResto(x int, y int) (resultado int, resto int) { // Possibilidade de retornar mais de 1 valor na função
	resultado = x / y
	resto = x % y
	return
}