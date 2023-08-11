package server

import (
	handlerUser "github.com/1rhino/clean_architecture/app/modules/users/handlers"
	repositoryUser "github.com/1rhino/clean_architecture/app/modules/users/repositories"
	userUseCase "github.com/1rhino/clean_architecture/app/modules/users/usecase"
)

func SetupRoutes(server *Server) {

	// Auth
	userRepo := repositoryUser.NewUserRepo(server.DB)
	userUseCase := userUseCase.NewUserUseCase(userRepo)
	userHandler := handlerUser.NewUserHandlers(userUseCase)

	api := server.Fiber.Group("/api/v1")

	user := api.Group("/users")
	user.Post("/signup", userHandler.SignUpUser())
}
