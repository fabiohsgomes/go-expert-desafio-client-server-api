package main

import (
	"context"
	"desfio-client-server-api/internal/consultacotacaoserver"
	"desfio-client-server-api/internal/infra/db/sqlite"
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	sqlite.Load()

	http.HandleFunc("GET /cotacao", ConsultaCotacao)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao subir o servidor: %s", err.Error())
	}
}

func ConsultaCotacao(w http.ResponseWriter, r *http.Request) {
	requestContext := r.Context()
	go checkingPerformContext(requestContext)

	cotacao, err := consultacotacaoserver.ProcessaCotacao(requestContext)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(cotacao)
	}
}

func checkingPerformContext(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
