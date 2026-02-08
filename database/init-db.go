package database

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kasir-app/config"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
)

var GormDB *gorm.DB
var SqlDB *sql.DB

func InitDB(connectionString string) {
	cfg, _ := pgx.ParseConfig(connectionString)

	cfg.StatementCacheCapacity = 0
	cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	db := stdlib.OpenDB(*cfg)

	// Best practice pool config
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Second)

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal("errors ping database ", err)
	}

	fmt.Println("Database connected successfully")
	SqlDB = db
}

func InitDBGorm() {

	//manual define
	fmt.Println(config.AppConfig.DbGormUser, " user nya")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.AppConfig.DbGormHost, config.AppConfig.DbGormUser, config.AppConfig.DbGormPass, config.AppConfig.DbGormDB)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	GormDB = db

	realValue := "anton gaming"
	var pointerValue *string
	pointerValue = &realValue
	log.Println(*pointerValue + " before")
	value, _ := testParamPointer(pointerValue)
	log.Println(value)
	log.Println(*pointerValue + " after")

	log.Println("✅ Database connected")
}

func testParamPointer(req *string) (string, error) {
	var response string
	response = *req + " yahuuuuu" // get real value from pointer

	return response, nil
}
