package web

import (
	"net/http"
	"task-gateway/web/handlers"
	"task-gateway/web/middlewares"
)

func InitRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle(
		"POST /auth/signup",
		manager.With(
			http.HandlerFunc(handlers.SignUp),
		),
	)
	mux.Handle(
		"POST /auth/signin",
		manager.With(
			http.HandlerFunc(handlers.SignIn),
		),
	)
	mux.Handle(
		"POST /auth/signout",
		manager.With(
			http.HandlerFunc(handlers.SignOut),
		),
	)
	mux.Handle(
		"GET /users/profile",
		manager.With(
			http.HandlerFunc(handlers.GetUserProfile),
		),
	)

	// mux.Handle(
	// 	"POST /users/verify-otp",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.VerifyOTP),
	// 	),
	// )
	// mux.Handle(
	// 	"POST /users/signin",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.SignInUser),
	// 	),
	// )
	// mux.Handle(
	// 	"GET /users",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.GetAllUsers),
	// 	),
	// )
	// mux.Handle(
	// 	"GET /users/",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.GetUserById),
	// 		middlewares.AuthMiddleware,
	// 	),
	// )
	// mux.Handle(
	// 	"PUT /users",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.UpdateUser),
	// 		middlewares.AuthMiddleware,
	// 	),
	// )
	// mux.Handle(
	// 	"DELETE /users",
	// 	manager.With(
	// 		http.HandlerFunc(handlers.DeleteUser),
	// 		middlewares.AuthMiddleware,
	// 	),
	// )

}
