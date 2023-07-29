package router

import (
	"task/internal/app/controller"
	"task/internal/app/repo"
	"task/internal/app/service"
	"task/pkg/db"

	"github.com/go-chi/chi/middleware"

	"github.com/go-chi/chi"
)

var (
	database      = db.GetDBConnection()
	statusRepo    = repo.NewCompanyRepo(database)
	statusService = service.NewStatusService(statusRepo)
	statusCtrl    = controller.NewStatusController(statusService)
)

// To set up router.
func Routes() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))
	r.Get("/company/{country}", statusCtrl.GetDetailsByCountry)
	return r
}
