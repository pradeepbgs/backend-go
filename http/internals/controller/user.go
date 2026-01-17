package controller

import (
	"net/http"
	"strconv"

	"github.com/pradeepbgs/internals/service"
	"github.com/pradeepbgs/internals/utils"
)

type UserController struct {
	service service.UserServiceInterface
}

func NewUserController(service service.UserServiceInterface) *UserController {
	return &UserController{service: service}
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.service.GetUsers()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.Json(w, http.StatusOK, users)
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := c.service.GetUserById(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "User not found")
		return
	}
	utils.Json(w, http.StatusOK, user)
}
