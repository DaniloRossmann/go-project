package main

import (
	"encoding/json"
	"fmt"
	"log"

	operation "github.com/DaniloRossmann/go-project/pkg/operations"
)

//criando metodos
type Veiculo struct {
	Nome  string `json:"nome"`
	Marca string `json:"marca"`
	Ano   int    `json:"ano"`
}

type Carro struct {
	Veiculo
	Qtdportas int `json:"qtd_porta"`
}

//o metodo
func (c Veiculo) andar() {
	fmt.Println(c.Nome, "andou")
}

func (c Carro) Exibe() {
	fmt.Printf("Nome: %s. Marca: %s. Ano: %d. Qtd Portas: %d \n", c.Nome, c.Marca, c.Ano, c.Qtdportas)
}

//--------------------------------------------------------
func main() {
	res, err := operation.Diff(2, 2)
	//valorSum := operation.Sum(15, 10)
	//vari := operation.Variable
	//fmt.Printf("%v \n %v ", valorSum, vari)

	//tratando erros
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println(res)
	}

	//retornando funcao dentro de funcao
	/*fmt.Println(operation.Result(54, 54, 54, 54)())*/

	//chamando metodo e atribuindo valor
	carro := Carro{
		Veiculo: Veiculo{
			Nome:  "Gol bolinha",
			Marca: "VW",
			Ano:   2010,
		},
		Qtdportas: 4,
	}

	//veiculo.andar()
	//carro.Exibe()

	//transformando struct em JSON
	veiculoJson, err := json.Marshal(carro)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(veiculoJson))

	//transformando JSON em struct
	carroJson2 := `{"nome":"Gol bolinha","marca":"VW","ano":2010,"qtd_porta":4}`
	carro2 := Carro{}
	json.Unmarshal([]byte(carroJson2), &carro2)

	fmt.Println(carro2)

}
