package main

import (
	"desfio-client-server-api/internal/consultacotacaoclient"
	"log"
	"os"
	"text/template"
)

func main() {
	cotacao, err := consultacotacaoclient.Consulta()
	if err != nil {
		log.Fatalf("Falha ao consultar a cotação :: %s", err.Error())

		return
	}

	logCotacao(cotacao)
}

func logCotacao(cotacao consultacotacaoclient.Cotacao) {
	templateBody := "Dólar: {{.Cotacao}}"
	tmpl, err := template.New("cotacao").Parse(templateBody)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl.Execute(file, cotacao)
}
