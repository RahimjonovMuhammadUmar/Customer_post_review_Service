package db

import (
	"fmt"
	"exam/customer_service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres drivers
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println("error while connecting to db", err)
		return nil, err
	}
	return connDB, nil
}


func ConnectToDBForSuiter(cfg config.Config) (*sqlx.DB, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDB, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		fmt.Println("error while connecting to db", err)
		return nil, func() {}
	}

	CleanUp := func ()  {
		connDB.Close()

	}

	return connDB, CleanUp
}
