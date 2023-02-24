//go:generate rm -fr mocks
//go:generate mockery --all

package db

import (
	"log"
	"os"
	"strconv"

	"github.com/getsentry/sentry-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GormMysql(dsn string) *gorm.DB {
	// logLevel (env: DB_LOG_MODE)
	// 1 = Silent (not printing any log)
	// 2 = Error (only printing in case of error)
	// 3 = Warn (print error + warning)
	// 4 = Info (print all type of log)
	logLevel, _ := strconv.Atoi(os.Getenv("DB_LOG_MODE"))
	if logLevel <= 0 {
		logLevel = 2
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.LogLevel(logLevel)),
	})
	if err != nil {
		log.Println("gorm.Open:", err)
		sentry.CaptureException(err)
	}
	return db
}
