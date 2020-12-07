package odoohelper

import (
	"database/sql"
	"fmt"
	"qibla-backend/pkg/functioncaller"
	"qibla-backend/pkg/logruslogger"
	"time"
)

type Connection struct {
	Host        string
	DbName      string
	User        string
	Password    string
	Port        string
	Location    *time.Location
	SslMode     string
	SslCert     string
	SslKey      string
	SslRootCert string
}

func (c Connection) DbConnect() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=UTC", c.User, c.Password, c.Host,c.Port, c.DbName, c.SslMode,
	)

	if c.SslMode == "require" {
		connStr = fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s&TimeZone=UTC&sslcert=%s&sslkey=%s&sslrootcert=%s",
			c.User, c.Password, c.Host, c.Port, c.DbName, c.SslMode, c.SslCert, c.SslKey, c.SslRootCert,
		)
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil{
		logruslogger.Log(logruslogger.WarnLevel,err.Error(),functioncaller.PrintFuncName(),"odoo-db-connection")
	}

	// migrations := &migrate.FileMigrationSource{
	// 	Dir: "../db/migrations",
	// }
	// n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	// if err != nil {
	// 	log.Fatal("Error migration := ", err.Error())
	// }
	// fmt.Printf("Applied %d migrations!\n", n)

	return db, err
}
