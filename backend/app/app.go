package app

import (
	"context"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	Config   Config
	DBClient *gorm.DB
	router   http.Handler
}

func Init(config Config) *App {
	app := &App{
		Config: config,
	}
	return app
}

func (app *App) Start(ctx context.Context) error {
	fmt.Println("Starting backend...")

	if err := app.connectDB(ctx); err != nil {
		return err
	}

	fmt.Println("Connected to the db!")

	app.router = loadRoutes(app)

	http_server := &http.Server{
		Addr:    ":8080",
		Handler: app.router,
	}

	fmt.Println("Server started!")

	http_ch := make(chan error, 1)
	go func() {
		err := http_server.ListenAndServe()
		if err != nil {
			http_ch <- fmt.Errorf("failed to start http server: %w", err)
		}
		close(http_ch)
	}()

	select {
	case err := <-http_ch:
		return err
	}
}

func (app *App) connectDB(ctx context.Context) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		app.Config.DatabaseConfig.User,
		app.Config.DatabaseConfig.Password,
		app.Config.DatabaseConfig.Addr,
		app.Config.DatabaseConfig.DBName,
	)

	gorm, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	app.DBClient = gorm
	return nil
}
