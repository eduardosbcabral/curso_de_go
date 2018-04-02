package operadora

import (
		"strconv"
		"curso_de_go/pacotes/prefixo" // ou "../prefixo"
)

//NomeOperadora representa o nome da operadora que está usando
var NomeOperadora = "Claro"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora