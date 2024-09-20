package main

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"paisleypark/kms/http/routes"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		zap.L().Fatal("Failed to read in config", zap.Error(err))
	}

	configureLogger()
	migrateDb(viper.GetString("CONNECTION_STRINGS__DB_CONNECTION"))

	r := gin.Default()

	configureRoutes(r)

	r.Run(":3003")
}

func configureRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/keys", routes.GETKeys)

	r.POST("/keys", routes.POSTKeys)
}

func configureLogger() {
	var logger *zap.Logger

	env := viper.GetString("APP_ENV")

	if env == "production" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	zap.ReplaceGlobals(logger)

	defer logger.Sync()

	logger.Info("Zap logger configured successfully",
		zap.String("environment", env))
}

func migrateDb(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Fatal(err.Error())
	}

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
		zap.L().Error("An error occured", zap.Error(err))
	}

	for _, file := range files {
		sqlScript, err := os.ReadFile(filepath.Join(migrationsPath, file.Name()))
		if err != nil {
			zap.L().Error("An error occured", zap.Error(err))
		}

		err = db.Exec(string(sqlScript)).Error
		if err != nil {
			zap.L().Error("An error occured", zap.Error(err))
		}
	}
}
