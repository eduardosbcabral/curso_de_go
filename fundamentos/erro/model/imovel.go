package model

import (
	"errors"
)

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

func (i *Imovel) Validate() (err error) {
	err = nil

	if i.valor < 50000 {
		return errors.New("O valor informado não é válido para um imóvel.")
	}

	if i.valor > 10000000 {
		return errors.New("O valor é muito alto. Corretor, por favor verifique sua avaliação.")
	}

	return 
}