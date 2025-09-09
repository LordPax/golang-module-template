package database

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	*core.Provider
	DB            *gorm.DB
	dotenvService *dotenv.DotenvService
}

func NewDatabaseService(module *DatabaseModule) *DatabaseService {
	return &DatabaseService{
		Provider:      core.NewProvider("DatabaseService"),
		dotenvService: module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (ds *DatabaseService) OnInit() error {
	fmt.Printf("Initializing %s\n", ds.GetName())
	return ds.Connect()
}

func (ds *DatabaseService) Connect() error {
	fmt.Println("Connecting to the database")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		ds.dotenvService.Get("DB_HOST"),
		ds.dotenvService.Get("DB_USER"),
		ds.dotenvService.Get("DB_PASSWORD"),
		ds.dotenvService.Get("DB_NAME"),
		ds.dotenvService.Get("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	ds.DB = db
	return nil
}

func (ds *DatabaseService) GetDB() *gorm.DB {
	return ds.DB
}

func (ds *DatabaseService) Close() error {
	sqlDB, err := ds.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (ds *DatabaseService) Migrate(models ...interface{}) error {
	return ds.DB.AutoMigrate(models...)
}

func (ds *DatabaseService) Table(name string) *gorm.DB {
	return ds.DB.Table(name)
}
