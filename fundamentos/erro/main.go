package main

import (
	"fmt"
	"encoding/json"
	"curso_de_go/fundamentos/erro/model"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Lucy"
	casa.X = 17
	casa.Y = 29
	casa.SetValor(100000000)
	err := casa.Validate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Casa é: %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)	
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}

	fmt.Println("A casa em JSON: ", string(objJSON))

}