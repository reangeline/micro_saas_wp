package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/reangeline/micro_saas/internal/presentation/controller"
)

func InitializeUserRoutes(controller *controller.UserController, r chi.Router) {

	r.Route("/users", func(r chi.Router) {
		r.Post("/", controller.CreateUserRest)
		r.Get("/", controller.FindAll)
		r.Get("/{email}", controller.FindUserByEmailRest)
		r.Put("/", controller.UpdateByEmail)
	})

}
