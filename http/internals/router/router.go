package router

import (
	"net/http"

	"github.com/pradeepbgs/internals/controller"
	"github.com/pradeepbgs/internals/repository"
	"github.com/pradeepbgs/internals/service"
)



func SetupRouter(router  *http.ServeMux) *http.ServeMux {
	
	// user routes register
	userRepo := repository.FakeUserRespository()
	userServ := service.NewUserService(userRepo)
	userctrl := controller.NewUserController(userServ)
	router.HandleFunc("/users",userctrl.GetUsers)
	router.HandleFunc("/user",userctrl.GetUserById)
	
	
	// in the last we return the main router
	return router
}