package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initialise la base de données MySQL
func InitDB(username, password, hostname, dbname string) (*sql.DB, error) {
	// Création de la chaîne de connexion à la base de données MySQL
	dataSourceName := username + ":" + password + "@tcp(" + hostname + ")/" + dbname

	// Connexion à la base de données MySQL
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture de la base de données :", err)
		return nil, err
	}

	// Vérification de la connexion à la base de données
	err = db.Ping()
	if err != nil {
		log.Fatal("Impossible de se connecter à la base de données :", err)
		return nil, err
	}

	// Crée la table si elle n'existe pas
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS computer_prices (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		site TEXT,
		price TEXT
	);`)
	if err != nil {
		log.Fatal("Erreur lors de la création de la table :", err)
		return nil, err
	}

	return db, nil
}
