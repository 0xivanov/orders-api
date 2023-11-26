package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router  http.Handler
	redisDb redis.Client
}

func New() *App {
	app := &App{
		redisDb: *redis.NewClient(&redis.Options{}),
	}
	app.loadRoutes()
	return app
}

func (app *App) Start(kuramiqnko context.Context) error {
	err := app.redisDb.Ping(kuramiqnko).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis db: %v", err)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	ch := make(chan error, 1)
	defer close(ch)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %v", err)
		}
	}()

	select {
	case err = <-ch:
		return err
	case <-kuramiqnko.Done():
		fmt.Println("aslkdfj")
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
