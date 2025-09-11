package user

import (
	"golang-api/core"
	"golang-api/log"
	"golang-api/media"
	"golang-api/middleware"
	"golang-api/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	*core.Provider
	userService     *UserService
	userMiddleware  *UserMiddleware
	queryService    *query.QueryService
	mediaMiddleware *media.MediaMiddleware
	logService      *log.LogService
}

func NewUserController(module *UserModule) *UserController {
	return &UserController{
		Provider:        core.NewProvider("UserController"),
		userService:     module.Get("UserService").(*UserService),
		userMiddleware:  module.Get("UserMiddleware").(*UserMiddleware),
		queryService:    module.Get("QueryService").(*query.QueryService),
		mediaMiddleware: module.Get("MediaMiddleware").(*media.MediaMiddleware),
		logService:      module.Get("LogService").(*log.LogService),
	}
}

func (uc *UserController) RegisterRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
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
//	@Summary		find all users
//	@Description	find all users
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]User
//	@Failure		500	{object}	utils.HttpError
//	@Router			/users/ [get]
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
func (uc *UserController) FindOne(c *gin.Context) {
	user := c.MustGet("user").(*User)
	c.JSON(http.StatusOK, user)
}

// FindMe godoc
//
//	@Summary		find connected user
//	@Description	find connected user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	User
//	@Router			/users/me [get]
func (uc *UserController) FindMe(c *gin.Context) {
	connectedUser, _ := c.MustGet("connectedUser").(*User)
	c.JSON(http.StatusOK, connectedUser)
}

// Update godoc
//
//	@Summary		update user
//	@Description	update user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		UpdateUserDto	true	"User"
//	@Param			id		path		string			true	"User ID"
//	@Success		200		{object}	User
//	@Failure		400		{object}	utils.HttpError
//	@Failure		401		{object}	utils.HttpError
//	@Failure		404		{object}	utils.HttpError
//	@Failure		500		{object}	utils.HttpError
//	@Router			/users/{id} [patch]
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
//	@Summary		update user image
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
//	@Router			/users/{user}/image [post]
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
//	@Summary		delete user
//	@Description	delete user
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"User ID"
//	@Success		204
//	@Failure		400	{object}	utils.HttpError
//	@Failure		401	{object}	utils.HttpError
//	@Failure		404	{object}	utils.HttpError
//	@Failure		500	{object}	utils.HttpError
//	@Router			/users/{id} [delete]
func (uc *UserController) Delete(c *gin.Context) {
	user := c.MustGet("user").(*User)
	uc.userService.Delete(user.ID)
	c.JSON(http.StatusNoContent, nil)
}
