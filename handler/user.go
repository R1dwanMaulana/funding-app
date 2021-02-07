package handler

import (
	"fundingapp/helper"
	"fundingapp/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	// struct diatas di passing sbg parameter service

	var input user.RegisterUserInput

	//pengecekan error

	err := c.ShouldBindJSON(&input)
	if err != nil {
		// better message error
		errors := helper.FormatValidationError(err)
		// map key: string value interface
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	// jika user error
	if err != nil {
		response := helper.APIResponse("Register Account Failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtservice.GenerateToken()
	formatter := user.FormatUser(newUser, "tokentokentokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// user mamasukkan input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input strruct
	// input struct passing service
	// di service mencari dg bantuan repository user dengan email ...
	// mencocokkan password

	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		// map key: string value interface
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "tokentokentoken")

	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
