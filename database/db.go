package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func ConexaoDB() {
	strConexao := "host=localhost user=root password=root dbname=personalidades port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(strConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados.")
	}
}
