package connection

import (
	"database/sql"
	// "encoding/csv"
	// "encoding/json"
	"fmt"
	"os"
	"strconv"

	// "strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/micro/go-micro/util/log"
	common "github.com/skinnyguy/spectra-services/core/model"

	// Postgres dialect
	_ "github.com/lib/pq"
)

type Connection interface {
	ExecuteDBQuery(qapi string) ([][]string, error)
	ExecuteDBInsert(qapi string, args ...interface{}) error
	DB() *sqlx.DB
}

type connection struct {
	cgf *common.SpectraConfig
	db  *sqlx.DB
}

func NewConnection(cfg *common.SpectraConfig) Connection {
	return &connection{
		cfg,
		initDatabaseConnection(cfg),
	}
}

// Method to Execute query DB ...
func (connection *connection) ExecuteDBQuery(qapi string) ([][]string, error) {
	var records [][]string

	rows, errConnection := connection.db.Query(qapi)

	if errConnection == nil {
		records = convertToCsv(rows)
	}

	if errConnection != nil {
		defer rows.Close()
	}

	return records, errConnection
}

// Method Insert To Database ExecuteDBInsert ...
func (connection *connection) ExecuteDBInsert(qapi string, args ...interface{}) error {
	_, errConnection := connection.db.Exec(qapi, args...)

	if errConnection != nil {
		log.Info("Failed to execute insert Query", errConnection)
	}

	return errConnection
}

// Method Insert To Call DB Instance ...
func (connection *connection) DB() *sqlx.DB {
	return connection.db
}

func convertToCsv(rows *sql.Rows) (results [][]string) {
	cols, err := rows.Columns()
	if err != nil {
		log.Info("Failed to get columns", err)
		return
	}

	rawResult := make([][]byte, len(cols))
	results = append(results, []string{"HEADER"})

	dest := make([]interface{}, len(cols))
	for i := range rawResult {
		dest[i] = &rawResult[i]
	}

	for rows.Next() {
		var result []string
		err = rows.Scan(dest...)
		if err != nil {
			log.Info("Failed to scan row", err)
			return
		}

		for _, raw := range rawResult {
			if raw == nil {
				result = append(result, "")
			} else {
				result = append(result, string(raw))
			}
		}

		results = append(results, result)
	}

	return results
}

func initDatabaseConnection(config *common.SpectraConfig) *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	timeout := os.Getenv("DB_TIMEOUT")
	idleConn, _ := strconv.Atoi(os.Getenv("DB_IDLE_CONNECTION"))
	openConn, _ := strconv.Atoi(os.Getenv("DB_OPEN_CONNECTION"))
	connectionLifeTime, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_LIFETIME"))

	connectionstring := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s connect_timeout=%s sslmode=disable",
		host, port, user, pass, dbname, timeout,
	)

	connectionDatabase, connectionError := sqlx.Open("postgres", connectionstring)
	if connectionError != nil {
		log.Errorf("Error establishing connection to %s database: %v", dbname, connectionError)
		panic(connectionError)
	}

	pingError := connectionDatabase.Ping()
	if pingError != nil {
		log.Errorf("Error connecting to %s database: %s", dbname, pingError)
		panic(pingError)
	}

	connectionDatabase.SetMaxIdleConns(idleConn)
	connectionDatabase.SetMaxOpenConns(openConn)
	connectionDatabase.SetConnMaxLifetime(time.Duration(connectionLifeTime) * time.Second)

	return connectionDatabase
}
