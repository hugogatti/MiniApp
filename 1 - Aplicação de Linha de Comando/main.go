package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")
	
	//LINHA DE COMANDO É NOME DO MODULO E APP É NOME DO PACKGE QUE CRIAMOS
	aplicacao := app.Gerar()
	//DOIS JEITOS PARA TRATAR O ERRO
	//erro := aplicacao.Run(os.Args)
	//if erro != nil {
	//	log.Fatal(erro)
	//}
	//OU
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}