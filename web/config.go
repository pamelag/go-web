package main

import (
	"os"

	"github.com/jackc/pgx"
)

const (
	defaultPort    = "8080"
	defaultPGURL   = "127.0.0.1"
	defaultPGPort  = "5432"
	defaultPGUser  = "postgres"
	defaultPGPass  = ""
	defaultDBName  = "rightprism"
	//defaultDBType  = "postgres"
	maxConnections = 50
)

var (
	port   = envString("PORT", defaultPort)
	dburl  = envString("DB_HOST", defaultPGURL)
	dbport = envString("DB_PORT", defaultPGPort)
	dbuser = envString("DB_USER", defaultPGUser)
	dbpass = envString("DB_PASS", defaultPGPass)
	dbname = envString("DB_NAME", defaultDBName)
	//dbtype = envString("DB_TYPE", defaultDBType)

	httpAddr = ":" + port
)

func createConnPool() (*pgx.ConnPool, error) {
	connConfig := pgx.ConnConfig{Host: dburl, User: dbuser, Password: dbpass, Database: dbname}
	config := pgx.ConnPoolConfig{ConnConfig: connConfig, MaxConnections: maxConnections}
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
