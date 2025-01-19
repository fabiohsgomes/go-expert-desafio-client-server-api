package consultacotacaoclient

import "desfio-client-server-api/internal/infra/client/server"

type Cotacao struct{
	Cotacao string `json:"cotacao"`
}

func Consulta() (cotacao Cotacao, err error) {
	contacaoResult, err:= server.ConsultaCotacao()
	if err != nil {
		return cotacao, err
	}

	cotacao = Cotacao{Cotacao: contacaoResult.Cotacao}

	return cotacao, err
}