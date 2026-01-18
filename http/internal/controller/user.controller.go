package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pradeepbgs/internal/service"
	"github.com/pradeepbgs/internal/utils"
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

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request payloads")
		return
	}

	if req.Name == "" || req.Email == "" {
		utils.Error(w, http.StatusBadRequest, "Name and email are required")
		return
	}
	
	createdUser, err := c.service.CreateUser(req.Name,req.Email)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "couldn't create the user: "+err.Error())
		return
	}
	
	utils.Json(w, 200,createdUser)
}
