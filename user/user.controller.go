package user

import (
	"golang-api/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*core.Provider
	userService    *UserService
	userMiddleware *UserMiddleware
}

func NewUserController(module *UserModule) *UserController {
	return &UserController{
		Provider:       core.NewProvider("UserController"),
		userService:    module.Get("UserService").(*UserService),
		userMiddleware: module.Get("UserMiddleware").(*UserMiddleware),
	}
}

func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	users.GET("/", uc.FindAll)
	users.GET("/:id",
		uc.userMiddleware.GetUser("id"),
		uc.FindByID,
	)
}

// FindAll godoc
//
//	@Summary		find all users
//	@Description	find all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]User
//	@Failure		500	{object}	utils.HttpError
//	@Router			/users/ [get]
func (uc *UserController) FindAll(c *gin.Context) {
	users, err := uc.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// FindByID godoc
//
//	@Summary		find users by id
//	@Description	find users by id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"User ID"
//	@Success		200	{object}	User
//	@Failure		404	{object}	utils.HttpError
//	@Failure		500	{object}	utils.HttpError
//	@Router			/users/{id} [get]
func (uc *UserController) FindByID(c *gin.Context) {
	user := c.MustGet("user").(*User)
	c.JSON(http.StatusOK, user)
}
