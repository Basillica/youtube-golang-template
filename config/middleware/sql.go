package middleware

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SQLMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db *gorm.DB
		var sqlDB *sql.DB
		var err error

		dns := ""
		if db, err = gorm.Open(sqlserver.Open(dns), &gorm.Config{}); err != nil {
			panic(err)
		}
		if sqlDB, err = db.DB(); err != nil {
			panic(err)
		}
		sqlDB.SetConnMaxLifetime(time.Minute)
		sqlDB.SetConnMaxIdleTime(100)
		c.Set("sqlClient", db)
	}
}
