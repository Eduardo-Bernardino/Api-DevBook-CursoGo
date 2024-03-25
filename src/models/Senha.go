package models

//Representa o formato da requisicao da alteração de senha
type Senha struct {
	Nova  string `json: "nova"`
	Atual string `json: "atual"`
}
