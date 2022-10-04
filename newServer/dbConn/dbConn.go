package dbconn

import (
	"log"
	"os"
	"time"

	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestLog(testStr string) {
	log.Println(testStr)
}

func GetDBConnector() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.Config{
		User:                 os.Getenv("USER_NAME"),
		Passwd:               os.Getenv("PASSWORD"),
		Net:                  "tcp",
		Addr:                 os.Getenv("HOST"),
		Collation:            "utf8mb4_general_ci",
		Loc:                  time.UTC,
		MaxAllowedPacket:     4 << 20.,
		AllowNativePasswords: true,
		CheckConnLiveness:    true,
		DBName:               os.Getenv("DB_NAME"),
	}
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		panic(err)
	}
	db := sql.OpenDB(connector)
	return db
}
