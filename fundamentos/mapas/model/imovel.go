package model

type Imovel struct {
	X 		int 	`json:"coordenada_x"`
	Y 		int 	`json:"coordenada_y"`
	Nome 	string	`json:"nome"`
	valor 	int
}

func (i *Imovel) SetValor(valor int) { // Função pertence a essa struct Imovel
	i.valor = valor
}

func (i *Imovel) GetValor() int {
	return i.valor
}