package auth

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"
	"golang-api/middleware"
	"golang-api/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	*core.Provider
	authService    *AuthService
	authMiddleware *AuthMiddleware
	userService    *user.UserService
	dotenvService  *dotenv.DotenvService
}

func NewAuthController(module *AuthModule) *AuthController {
	return &AuthController{
		Provider:       core.NewProvider("AuthController"),
		authService:    module.Get("AuthService").(*AuthService),
		authMiddleware: module.Get("AuthMiddleware").(*AuthMiddleware),
		userService:    module.Get("UserService").(*user.UserService),
		dotenvService:  module.Get("DotenvService").(*dotenv.DotenvService),
	}
}

func (ac *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")
	auth.POST("/login",
		middleware.Validate[LoginDto](),
		ac.Login,
	)
	auth.POST("/register",
		middleware.Validate[user.CreateUserDto](),
		ac.Register,
	)
}

func (ac *AuthController) setAuthCookies(c *gin.Context, token *Token) {
	cookieSecure, _ := strconv.ParseBool(ac.dotenvService.Get("COOKIE_SECURE"))
	c.SetCookie("access_token", token.AccessToken, 3600, "/", "", cookieSecure, true)
	c.SetCookie("refresh_token", token.RefreshToken, 3600*24*30, "/", "", cookieSecure, true)
}

func (ac *AuthController) clearAuthCookies(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		LoginDto	true	"Login user"
//	@Success		200		{object}	LoginSuccessResponse
//	@Failure		400		{object}	utils.HttpError
//	@Router			/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	body, _ := c.MustGet("body").(LoginDto)

	user, err := ac.userService.FindOneBy("email", body.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !user.Verified {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account not verified"})
		return
	}

	if !user.ComparePassword(body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if ac.authService.DeleteTokensByUserID(user.ID) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tokens"})
		return
	}

	token := &Token{UserID: user.ID}

	jwtKey := ac.dotenvService.Get("JWT_SECRET_KEY")
	if token.GenerateTokens(jwtKey) != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	if err := ac.authService.CreateToken(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	ac.setAuthCookies(c, token)
	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken, "refresh_token": token.RefreshToken})
}

// Register godoc
//
//	@Summary		Register user
//	@Description	Register user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			register	body	user.CreateUserDto	true	"Register user"
//	@Success		201
//	@Failure		400	{object}	utils.HttpError
//	@Router			/auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	body, _ := c.MustGet("body").(user.CreateUserDto)

	if ac.userService.IsUserExists(body.Email, body.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Username already exists"})
		return
	}

	defaultAvatarUrl := fmt.Sprintf("https://api.dicebear.com/9.x/initials/svg?seed=%s", body.Firstname+body.Lastname)

	user := user.User{
		ID:        uuid.New().String(),
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Username:  body.Username,
		Email:     body.Email,
		Profile:   defaultAvatarUrl,
		Roles:     []string{user.ROLE_USER},
	}

	if err := user.HashPassword(body.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := ac.userService.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// sendWelcomeAndVerificationEmails(user, c)
	c.Status(http.StatusCreated)
}
