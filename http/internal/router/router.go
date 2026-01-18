package router

import (
	"net/http"
	"github.com/pradeepbgs/internal/controller"
	sqlc "github.com/pradeepbgs/internal/db"
	"github.com/pradeepbgs/internal/repository"
	"github.com/pradeepbgs/internal/service"
)



func SetupRouter(router  *http.ServeMux, queries  *sqlc.Queries) *http.ServeMux {
	
	// user routes register
	userRepo := repository.NewUserRespository(queries)
	userServ := service.NewUserService(userRepo)
	userctrl := controller.NewUserController(userServ)
	router.HandleFunc("/users",userctrl.GetUsers)
	router.HandleFunc("/user/{id}",userctrl.GetUserById)
	router.HandleFunc("/createUser", userctrl.CreateUser)
	// in the last we return the main router
	return router
}