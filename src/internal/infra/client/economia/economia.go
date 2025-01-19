package economia

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

type EconomiaDTO struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func ConsultaCotacao(ctx context.Context) (economia EconomiaDTO, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return economia, err
	}

	response, err := client.Send(request)
	if err != nil {
		contextFail := ctx.Err()
		if contextFail != nil {
			message := "API economia excedeu o limite de 200ms"

			log.Println(message)

			return economia, errors.New(message)
		}

		return economia, err
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return economia, fmt.Errorf("erro na chamada da api: %d :: %s", response.StatusCode, response.Status)
	}

	json.Unmarshal(response.Body, &economia)

	return economia, err
}
