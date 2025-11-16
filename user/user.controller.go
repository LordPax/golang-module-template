package user

import (
	"fmt"
	ginM "golang-api/gin"
	"golang-api/log"
	"golang-api/media"
	"golang-api/middleware"
	"golang-api/query"
	"net/http"

	"github.com/LordPax/godular/core"
	"github.com/gin-gonic/gin"
)

type IUserController interface {
	core.IProvider
	RegisterRoutes()
	FindAll(c *gin.Context)
	FindOne(c *gin.Context)
	FindMe(c *gin.Context)
	Update(c *gin.Context)
	UploadImage(c *gin.Context)
	Delete(c *gin.Context)
}

type UserController struct {
	*core.Provider
	userService     IUserService
	userMiddleware  IUserMiddleware
	queryService    query.IQueryService
	mediaMiddleware media.IMediaMiddleware
	logService      log.ILogService
	ginService      ginM.IGinService
}

func NewUserController(module core.IModule) *UserController {
	return &UserController{
		Provider:        core.NewProvider("UserController"),
		userService:     module.Get("UserService").(IUserService),
		userMiddleware:  module.Get("UserMiddleware").(IUserMiddleware),
		queryService:    module.Get("QueryService").(query.IQueryService),
		mediaMiddleware: module.Get("MediaMiddleware").(media.IMediaMiddleware),
		logService:      module.Get("LogService").(log.ILogService),
		ginService:      module.Get("GinService").(ginM.IGinService),
	}
}

func (uc *UserController) OnInit() error {
	uc.RegisterRoutes()
	return nil
}

func (uc *UserController) RegisterRoutes() {
	fmt.Println("Registering User routes")
	users := uc.ginService.GetGroup().Group("/users")
	users.GET("/",
		uc.userMiddleware.IsLoggedIn(true),
		uc.queryService.QueryFilter(),
		uc.FindAll,
	)
	users.GET("/:user",
		uc.userMiddleware.IsLoggedIn(true),
		uc.userMiddleware.FindOne("user"),
		uc.FindOne,
	)
	users.POST("/:user/image",
		uc.userMiddleware.IsLoggedIn(true),
		uc.userMiddleware.FindOne("user"),
		uc.userMiddleware.IsMe(),
		uc.mediaMiddleware.FileUploader(media.IMAGE, media.SIZE_10MB, "profile"),
		uc.UploadImage,
	)
	users.PATCH("/:user",
		uc.userMiddleware.IsLoggedIn(true),
		uc.userMiddleware.FindOne("user"),
		uc.userMiddleware.IsMe(),
		middleware.Validate[UpdateUserDto](),
		uc.Update,
	)
	users.DELETE("/:user",
		uc.userMiddleware.IsLoggedIn(true),
		uc.userMiddleware.FindOne("user"),
		uc.userMiddleware.IsMe(),
		uc.Delete,
	)
	users.GET("/me",
		uc.userMiddleware.IsLoggedIn(true),
		uc.FindMe,
	)
}

// FindAll godoc
//
//	@Summary		Find all users
//	@Description	find all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]User
//	@Failure		500	{object}	utils.HttpError
//	@Router			/api/users/ [get]
func (uc *UserController) FindAll(c *gin.Context) {
	query, _ := c.MustGet("query").(query.QueryFilter)
	users, err := uc.userService.FindAll(query)
	tags := []string{"UserController", "FindAll"}
	if err != nil {
		uc.logService.Errorf(tags, "%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// FindByID godoc
//
//	@Summary		Find users by id
//	@Description	find users by id
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	path		string	true	"User ID"
//	@Success		200	{object}	User
//	@Failure		404	{object}	utils.HttpError
//	@Failure		500	{object}	utils.HttpError
//	@Router			/api/users/{user} [get]
func (uc *UserController) FindOne(c *gin.Context) {
	user := c.MustGet("user").(*User)
	user.Secure()
	c.JSON(http.StatusOK, user)
}

// FindMe godoc
//
//	@Summary		Find connected user
//	@Description	find connected user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Router			/api/users/me [get]
func (uc *UserController) FindMe(c *gin.Context) {
	connectedUser, _ := c.MustGet("connectedUser").(*User)
	connectedUser.Secure()
	c.JSON(http.StatusOK, connectedUser)
}

// Update godoc
//
//	@Summary		Update user
//	@Description	update user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body		UpdateUserDto	true	"User"
//	@Param			user		path		string			true	"User ID"
//	@Success		200		{object}	User
//	@Failure		400		{object}	utils.HttpError
//	@Failure		401		{object}	utils.HttpError
//	@Failure		404		{object}	utils.HttpError
//	@Failure		500		{object}	utils.HttpError
//	@Router			/api/users/{user} [patch]
func (uc *UserController) Update(c *gin.Context) {
	user, _ := c.MustGet("user").(*User)
	body, _ := c.MustGet("body").(UpdateUserDto)
	tags := []string{"UserController", "Update"}

	if uc.userService.IsUserExists(body.Email, body.Username) {
		uc.logService.Errorf(tags, "Email or username already in use")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or username already in use"})
		return
	}

	if body.Firstname != "" {
		user.Firstname = body.Firstname
	}
	if body.Lastname != "" {
		user.Lastname = body.Lastname
	}
	if body.Username != "" {
		user.Username = body.Username
	}
	if body.Email != "" {
		user.Email = body.Email
	}

	if err := uc.userService.Update(user); err != nil {
		uc.logService.Errorf(tags, "%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUserImage godoc
//
//	@Summary		Update user image
//	@Description	update user image
//	@Tags			user
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			user		path	string		true	"User ID"
//	@Param			upload[]		formData	file	true	"Image"
//	@Success		200			{object}	User
//	@Failure		400			{object}	utils.HttpError
//	@Failure		401			{object}	utils.HttpError
//	@Failure		404			{object}	utils.HttpError
//	@Failure		500			{object}	utils.HttpError
//	@Router			/api/users/{user}/image [post]
func (uc *UserController) UploadImage(c *gin.Context) {
	user, _ := c.MustGet("user").(*User)
	medias, _ := c.MustGet("medias").([]*media.Media)
	tags := []string{"UserController", "UpdateImage"}

	user.Profile = medias[0].Url

	if err := uc.userService.Update(user); err != nil {
		uc.logService.Errorf(tags, "%s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
//
//	@Summary		Delete user
//	@Description	delete user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	path	string	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	utils.HttpError
//	@Failure		401	{object}	utils.HttpError
//	@Failure		404	{object}	utils.HttpError
//	@Failure		500	{object}	utils.HttpError
//	@Router			/api/users/{user} [delete]
func (uc *UserController) Delete(c *gin.Context) {
	user := c.MustGet("user").(*User)
	uc.userService.Delete(user.ID)
	c.JSON(http.StatusNoContent, nil)
}
