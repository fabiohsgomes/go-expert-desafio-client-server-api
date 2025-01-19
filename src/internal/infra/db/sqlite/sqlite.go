package sqlite

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const databaseDir string = "./database"
const databaseFileName string = "cotacao.db"

func Load() {
	loadInfra()
}

func loadInfra() {
	criaDiretorio()
	criaDataBase()

	connection := getConnection()

	createTables(connection)
}

func GetConnection() *sql.DB  {
	return getConnection()
}

func getConnection() (db *sql.DB) {
	pathDatabase := filepath.Join(databaseDir, databaseFileName)
	db, _ = sql.Open("sqlite3", pathDatabase)

	return db
}

func createTables(db *sql.DB) {
	defer db.Close()

	script := `
	CREATE TABLE IF NOT EXISTS cotacao(
		idCotacao INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		valor REAL NOT NULL,
		data_criacao TEXT NOT NULL
	)`

	statement, err := db.Prepare(script)
	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
}

func criaDiretorio() {
	if _, err := os.Stat(databaseDir); os.IsNotExist(err) {
		os.Mkdir(databaseDir, 0777)
	}
}

func criaDataBase() {
	pathDatabase := filepath.Join(databaseDir, databaseFileName)

	if _, err := os.Stat(pathDatabase); os.IsNotExist(err) {
		if file, err := os.Create(pathDatabase); os.IsNotExist(err) {
			criaDiretorio()
			criaDataBase()
			return
		} else {
			log.Println("Base de dados configurada com sucesso")
			file.Close()
		}
	}
}
