package config

import (
	"github.com/joho/godotenv"
	"io/fs"
	"log"
	"os"
)

type Config struct {
	HostDB        string
	PortDB        string
	UserDB        string
	PasswordDB    string
	NameDB        string
	SchemeDB      string
	MigrationPath fs.FS
	MaxOpenDBConn string
	MaxIdleDBConn string
	MaxDBLifeTime string
	HTTPPort      string
}

func NewConfig() Config {
	var conf Config
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	conf.HostDB = os.Getenv("HOST_DB")
	conf.PortDB = os.Getenv("PORT_DB")
	conf.UserDB = os.Getenv("USER_DB")
	conf.PasswordDB = os.Getenv("PASSWORD_DB")
	conf.NameDB = os.Getenv("NAME_DB")
	conf.SchemeDB = os.Getenv("SCHEME_DB_TABLE")
	migrationPath := os.Getenv("MIGRATION_PATH")
	conf.MigrationPath = os.DirFS(migrationPath)
	conf.MaxOpenDBConn = os.Getenv("MAX_OPEN_DB_CONNECTION")
	conf.MaxIdleDBConn = os.Getenv("MAX_IDLE_DB_CONNECTION")
	conf.MaxDBLifeTime = os.Getenv("MAX_DB_LIFE_TIME")
	conf.HTTPPort = os.Getenv("HTTP_PORT")
	return conf

}
