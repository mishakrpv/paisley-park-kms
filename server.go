package main

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"paisleypark/kms/http/routes"
	config "paisleypark/kms/interfaces/configuration"
)

func main() {
	configureLogging()
	config.Config = NewConfigurationManager()

	connectionString := config.Config.Get("CONNECTION_STRINGS__DB_CONNECTION")

	err := migrateDb(connectionString)
	if err != nil {
		zap.L().Warn("Something went wrong while migrating the db",
		zap.String("connection_string", connectionString),
		zap.Error(err))
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	authorize := r.Group("/")

	authorize.Use(RequireAuthorization())
	{
		authorize.GET("/keys", routes.GETKeys)

		authorize.POST("/keys", routes.POSTKeys)
		authorize.POST("/encrypt", routes.POSTEncrypt)
		authorize.POST("/decrypt", routes.POSTDecrypt)
	}

	r.Run(":3003")
}

func configureLogging() {
	var logger *zap.Logger

	env := os.Getenv("ENV")

	if env == "" {
		env = "development"
		os.Setenv("ENV", env)
	}

	if env == "production" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)

	defer logger.Sync()

	logger.Info("Starting up", zap.String("environment", env))
}

func migrateDb(dsn string) (err error) {
	routes.Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal("Failed to open database connection",
			zap.String("dsn", dsn),
			zap.Error(err))
	}

	// TODO: debug the rest of the function
	migrationsPath := "infrastructure/migrations"

	var files []fs.FileInfo
	err = filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fileInfo, err := d.Info()
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() && filepath.Ext(d.Name()) == ".sql" {
			files = append(files, fileInfo)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for _, file := range files {
		sqlScript, err := os.ReadFile(filepath.Join(migrationsPath, file.Name()))
		if err != nil {
			return err
		}

		err = routes.Db.Exec(string(sqlScript)).Error
		if err != nil {
			return err
		}

		zap.L().Info("Migration has been applied", zap.String("from_filename", file.Name()))
	}

	return nil
}
