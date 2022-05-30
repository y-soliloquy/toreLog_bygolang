package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"torelog_bygolang/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// const (
// 	tableNameUser        = "users"
// 	tableNameTrainingLog = "trainingLogs"
// 	tableNameSession     = "sessions"
// )

func init() {

	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Fatalln(err)
	}

	// sqlite向けの処理
	// Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	name STRING,
	// 	email STRING,
	// 	password STRING,
	// 	created_at DATETIME)`, tableNameUser)

	// Db.Exec(cmdU)

	// cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	content TEXT,
	// 	satisfaction STRING,
	// 	weather STRING,
	// 	user_id INTEGER,
	// 	created_at DATETIME)`, tableNameTrainingLog)

	// Db.Exec(cmdT)

	// cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
	// 	id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 	uuid STRING NOT NULL UNIQUE,
	// 	email STRING,
	// 	user_id INTEGER,
	// 	created_at DATETIME)`, tableNameSession)

	// Db.Exec(cmdS)

}

// UUID生成
func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

// ハッシュ化されたパスワード生成
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
