package application

import (
	"net/http"

	"github.com/0xivanov/orders-api/handler"
	"github.com/0xivanov/orders-api/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("bazik ataka"))
	})
	router.Route("/orders", app.loadOrderRoutes)
	app.router = router
}

func (app *App) loadOrderRoutes(router chi.Router) {
	orderRoutes := &handler.Order{
		Repo: &repository.RedisRepo{
			Client: &app.redisDb,
		},
	}

	router.Get("/", orderRoutes.List)
	router.Get("/{id}", orderRoutes.GetById)
	router.Post("/", orderRoutes.Create)
	router.Put("/{id}", orderRoutes.UpdateById)
	router.Delete("/{id}", orderRoutes.DeleteById)

}
