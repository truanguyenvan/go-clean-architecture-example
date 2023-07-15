package gorm

import (
	"database/sql"
	"errors"
	"go-clean-architecture-example/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

// Gorm defines a interface for access the database.
type Gorm interface {
	DB() *gorm.DB
	SqlDB() *sql.DB
	Transaction(fc func(tx *gorm.DB) error) (err error)
	Close() error
	DropTableIfExists(value interface{}) error
}

// Config GORM Config
type Config struct {
	Debug           bool
	DBType          string
	DSN             string
	MaxLifetime     int
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime int
	ConnMaxIdleTime int
	TablePrefix     string
}

// _gorm gorm struct
type _gorm struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

// New Create gorm.DB and  instance
func New(c Config, logger logger.Logger) (Gorm, error) {
	var dial gorm.Dialector

	switch strings.ToLower(c.DBType) {
	case "mysql":
		dial = mysql.Open(c.DSN)
	case "postgres":
		dial = postgres.Open(c.DSN)
	default:
		return nil, errors.New("DBType does not support")
	}

	gConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
		Logger: logger,
	}

	db, err := gorm.Open(dial, gConfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if c.MaxLifetime != 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	}
	if c.ConnMaxLifetime != 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Second)
	}
	if c.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.ConnMaxIdleTime != 0 {
		sqlDB.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Second)
	}
	return &_gorm{
		db:    db,
		sqlDB: sqlDB,
	}, nil
}
