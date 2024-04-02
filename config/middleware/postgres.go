package middleware

import (
	"database/sql"
	"fmt"
	"time"

	"go-template/types/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func PostgresMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ae := c.MustGet("appConfig").(*config.AppConfig)

		connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s",
			ae.POSTGRES_USER, ae.POSTGRES_PASS, ae.POSTGRES_DBNAME, ae.POSTGRES_HOST, ae.POSTGRES_SSLMODE,
		)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		err = db.Ping()
		if err != nil {
			panic(err)
		}

		db.SetMaxIdleConns(100)
		db.SetConnMaxLifetime(time.Hour)

		c.Set("postgres", db)
	}
}
