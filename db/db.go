package db

import (
	"fmt"
	"log"
	"miniTwitter/configs"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/lib/pq"
)

// Initialize database connection then connect with postgres
func Init(config *configs.Configuration) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDatabase)
	m, err := migrate.New("file://db/migrations", dbURL)
	if err != nil {
		log.Fatal("error in creating migrations: ", zap.Error(err))
	}

	if err = m.Up(); err != nil {
		log.Println("error updating migrations: ", zap.Error(err))
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logger.Error,    // For debugging you can set Info level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,           // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	return db, nil
	
}
