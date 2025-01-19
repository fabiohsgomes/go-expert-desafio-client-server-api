package consultacotacaoserver

import (
	"context"
	"database/sql"
	"desfio-client-server-api/internal/infra/client/economia"
	"errors"
	"log"
	"time"
)

type Cotacao struct {
	Valor       string `json:"cotacao"`
	DataCriacao string `json:"-"`
}

func NewCotacao() Cotacao {
	return Cotacao{}
}

func (c *Cotacao) consulta(ctx context.Context) (err error) {
	economia, err := economia.ConsultaCotacao(ctx)
	if err != nil {
		return err
	}

	c.Valor = economia.Usdbrl.Bid
	c.DataCriacao = economia.Usdbrl.CreateDate

	return err
}

func (c *Cotacao) salva(ctx context.Context, dbConnection *sql.DB) (err error) {
	defer dbConnection.Close()

	dbConnection.SetMaxOpenConns(1)

	ctx, cancel := context.WithTimeout(ctx, time.Microsecond*10)
	defer cancel()

	script := `INSERT INTO cotacao(valor, data_criacao) VALUES(?, ?)`

	_, err = dbConnection.ExecContext(ctx, script, c.Valor, c.DataCriacao)
	if err != nil {
		contextFail := ctx.Err()
		if contextFail != nil {
			message := "não foi possível salvar, pois a operação excedeu o limite de 10ms"

			log.Println(message)

			return errors.New(message)
		}
	}

	return err
}
