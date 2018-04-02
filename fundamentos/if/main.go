package main

import "fmt"

func main() {

	meses := 0
	situacao := true
	cidade := "Brasília"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	} else {
		fmt.Println("Esse credor deve há muito tempo")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele está adimplente")
	}

	if cidade == "Brasília" {
		fmt.Println("O cliente vive no Distrito Federal")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status { // Essas variáveis (descricao e status) só funcionam no escopo do IF
		fmt.Println("Qual a situação do cliente? ", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")

}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}
	descricao = "O cliente está com o carnê em dia."
	return
}