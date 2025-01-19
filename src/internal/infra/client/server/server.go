package server

import (
	"context"
	"desfio-client-server-api/internal/infra/client"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CotacaoDTO struct {
	Cotacao string `json:"cotacao"`
}

func ConsultaCotacao() (cotacao CotacaoDTO, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		return cotacao, err
	}

	response, err := client.Send(request)
	if err != nil {
		contextFail := ctx.Err()
		if contextFail != nil {
			message := "API cotacao excedeu o limite de 300ms para retornar a consulta"
			log.Println(message)

			return cotacao, errors.New(message)
		}

		return cotacao, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return cotacao, fmt.Errorf("erro na chamada da api: %d :: %s", response.StatusCode, response.Status)
	}

	json.Unmarshal(response.Body, &cotacao)

	return cotacao, err
}
