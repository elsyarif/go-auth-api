package handler

import (
	"context"
	"github.com/elsyarif/go-auth-api/internal/applications/usecases"
	"github.com/elsyarif/go-auth-api/internal/domain/entities"
	"github.com/elsyarif/go-auth-api/pkg/common"
	"github.com/elsyarif/go-auth-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewUserHandler(addUser usecases.UserUseCase) UserHandler {
	return UserHandler{userUseCase: addUser}
}

func (h *UserHandler) Routes(app *gin.Engine) {
	user := app.Group("/users")
	user.POST("", h.PostUserHandler)
}

func (h *UserHandler) PostUserHandler(c *gin.Context) {
	ctx := context.Background()

	user := entities.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		appError := common.NewError(err, common.ValidationError)
		c.Error(appError)
		return
	}

	result, err := h.userUseCase.AddUser(ctx, user)
	if err != nil {
		c.Error(common.NewError(err, common.ResourceAlreadyExists))
		return
	}
	ss := entities.UserToResponse(result)
	c.JSON(http.StatusCreated, helper.ResponseJSON.Success("success", ss))
}
