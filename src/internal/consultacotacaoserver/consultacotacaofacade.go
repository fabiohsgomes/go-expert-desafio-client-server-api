package consultacotacaoserver

import (
	"context"
	"desfio-client-server-api/internal/infra/db/sqlite"
)

func ProcessaCotacao(ctx context.Context) (cotacaoDTO Cotacao, err error) {
	cotacao := NewCotacao()
	
	err = cotacao.consulta(ctx)
	if err != nil {
		return cotacaoDTO, err
	}
	
	err = cotacao.salva(ctx, sqlite.GetConnection())

	return cotacao, err
}
