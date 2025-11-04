package postgresql

import (
	"time"

	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
	"log"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	DBName   string
}

type PostgresDB struct {
	config Config
	db     *sql.DB
}

//TODO : idor for pk:id

func New(config Config) *PostgresDB {
	//connStr := "postgres://userservice:pass@localhost:5433/mydb?sslmode=disable"
	db, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.SetMaxOpenConns(25)                 // max number of open connections
	db.SetMaxIdleConns(25)                 // max number of idle connections
	db.SetConnMaxLifetime(5 * time.Minute) // how long a connection can be reused

	fmt.Println("Connected to Postgres successfully 🎉")
	return &PostgresDB{db: db, config: config}
}
