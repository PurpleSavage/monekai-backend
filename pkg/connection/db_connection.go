package connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// NewClient inicializa la conexión a GORM de forma genérica
func NewClient(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configuración del pool de conexiones
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}