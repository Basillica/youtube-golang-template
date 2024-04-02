package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-template/config/helpers"
	"go-template/config/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type API struct {
	Engine *gin.Engine
	srv    *http.Server
}

func (a *API) Init(version string) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	if helpers.GetEnv("APP_MODE", "") == "release" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Print("Running on release mode")
	} else {
		fmt.Println("Running on debug mode")
		if err := godotenv.Load("app.env"); err != nil {
			return nil, err
		}
	}

	a.Engine = r
	r.Use(a.WithMiddlewares(version)...)
	a.AddRoutes(version)
	a.WithServer()
	if !helpers.GetEnvAsBool("TEST_MODE", false) {
		a.MakeMigrationsSQL()
		a.MakeMigrationsPostgres()
	}
	return r, nil
}

func (a *API) WithMiddlewares(version string) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.ConfigMiddleware(),
		middleware.CORSMiddleware(),
	}
}

func (a *API) AddRoutes(version string) {
	AddUserRoutes(a.Engine, version)
	AddAccountsRoutes(a.Engine, version)
	AddMiscRoutes(a.Engine, version)
	AddSwaggerRoutes(a.Engine, version)
}

func (a *API) MakeMigrationsSQL() {
	fmt.Println("making migrations to sql database ...")
}

func (a *API) MakeMigrationsPostgres() {
	fmt.Println("making migrations to sql database ...")
}

func (a *API) WithServer() API {
	ch := make(chan API)
	go func() {
		a.srv = &http.Server{
			Addr:              fmt.Sprintf(":%s", "8080"),
			Handler:           a.Engine,
			ReadHeaderTimeout: 100,
		}
		ch <- *a
	}()
	return <-ch
}

func (a *API) Start(r *gin.Engine) {
	// build and start server
	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be caugth, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Panicln("Server Shutdown: ...", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("Server exiting ...")
}

func New(apiVersion string) API {
	api := API{}
	r, err := api.Init(apiVersion)
	if err != nil {
		panic(err)
	}
	api.Engine = r
	return api
}
