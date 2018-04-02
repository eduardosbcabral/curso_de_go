package matematica

func Calculo(funcao func(int, int) int, x int, y int) int {
	return funcao(x, y)
}