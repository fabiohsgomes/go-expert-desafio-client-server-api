# Observação

O projeto encontra-se na pasta **src**. Estando na pasta src, executar os comandos:
```
go mod tidy

#Executando o server 
go run cmd/server/main.go

#Executando o client
go run cmd/client/main.go
```

## Descrição do desafio

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:

- Servidor:

O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL 
e em seguida deverá retornar no formato JSON o resultado para o cliente apenas o valor atual do câmbio (campo "bid" do JSON). (feito)

O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080. (feito)

O server.go deverá registrar no banco de dados SQLite cada cotação recebida.
    - Criar banco de dados (feito)
    - Criar tabela (feito)
    - Criar código para registrar os dados. (feito)

Usando o package "context", o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
    - Timeout consulta api 200ms. (feito)
    - Timeout gravação no banco 10ms. (feito)

- Cliente:

O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar. (feito)
  
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). (feito)

O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}. (feito)

Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go. (feito)

- Contexto

Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente. (feito)
 
Ao finalizar, envie o link do repositório para correção.