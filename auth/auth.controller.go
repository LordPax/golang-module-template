package auth

import (
	"fmt"
	"golang-api/core"
	"golang-api/dotenv"
	ginM "golang-api/gin"
	"golang-api/log"
	"golang-api/middleware"
	"golang-api/token"
	"golang-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	*core.Provider
	authService    *AuthService
	tokenService   *token.TokenService
	userService    *user.UserService
	dotenvService  *dotenv.DotenvService
	userMiddleware *user.UserMiddleware
	logService     *log.LogService
	ginService     *ginM.GinService
}

func NewAuthController(module *AuthModule) *AuthController {
	return &AuthController{
		Provider:       core.NewProvider("AuthController"),
		authService:    module.Get("AuthService").(*AuthService),
		tokenService:   module.Get("TokenService").(*token.TokenService),
		userService:    module.Get("UserService").(*user.UserService),
		dotenvService:  module.Get("DotenvService").(*dotenv.DotenvService),
		userMiddleware: module.Get("UserMiddleware").(*user.UserMiddleware),
		logService:     module.Get("LogService").(*log.LogService),
		ginService:     module.Get("GinService").(*ginM.GinService),
	}
}

func (ac *AuthController) OnInit() error {
	ac.RegisterRoutes()
	return nil
}

func (ac *AuthController) RegisterRoutes() {
	fmt.Println("Registering Auth routes")
	auth := ac.ginService.Group.Group("/auth")
	auth.POST("/login",
		middleware.Validate[user.LoginDto](),
		ac.Login,
	)
	auth.POST("/register",
		middleware.Validate[user.CreateUserDto](),
		ac.Register,
	)
	auth.POST("/logout",
		ac.userMiddleware.IsLoggedIn(true),
		ac.Logout,
	)
	auth.POST("/refresh", ac.Refresh)
}

// Login godoc
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			login	body		user.LoginDto	true	"Login user"
//	@Success		200		{object}	token.LoginSuccessResponse
//	@Failure		400		{object}	utils.HttpError
//	@Router			/api/auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	body, _ := c.MustGet("body").(user.LoginDto)
	tags := []string{"AuthController", "Login"}

	user, err := ac.userService.FindOneBy("email", body.Email)
	if err != nil {
		ac.logService.Errorf(tags, "Invalid credentials: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if !user.Verified {
		ac.logService.Errorf(tags, "Account not verified for user ID: %s", user.ID)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account not verified"})
		return
	}

	if !user.ComparePassword(body.Password) {
		ac.logService.Errorf(tags, "Invalid credentials for user ID: %s", user.ID)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if ac.tokenService.DeleteByUserID(user.ID) != nil {
		ac.logService.Errorf(tags, "Failed to delete existing tokens for user ID: %s", user.ID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tokens"})
		return
	}

	token := &token.Token{UserID: user.ID}

	jwtKey := ac.dotenvService.Get("JWT_SECRET_KEY")
	if token.GenerateTokens(jwtKey) != nil {
		ac.logService.Errorf(tags, "Failed to generate tokens for user ID: %s", user.ID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	if err := ac.tokenService.Create(token); err != nil {
		ac.logService.Errorf(tags, "Failed to create token for user ID: %s, error: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	ac.authService.SetAuthCookies(c, token)
	ac.logService.Printf(tags, "User ID %s logged in successfully", user.ID)
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
//	@Router			/api/auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	body, _ := c.MustGet("body").(user.CreateUserDto)
	tags := []string{"AuthController", "Register"}

	if ac.userService.IsUserExists(body.Email, body.Username) {
		ac.logService.Errorf(tags, "Email or Username already exists: %s, %s", body.Email, body.Username)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or Username already exists"})
		return
	}

	user := &user.User{
		Firstname: body.Firstname,
		Lastname:  body.Lastname,
		Username:  body.Username,
		Email:     body.Email,
		Profile:   fmt.Sprintf("https://api.dicebear.com/9.x/initials/svg?seed=%s", body.Firstname+body.Lastname),
		Roles:     []string{user.ROLE_USER},
	}

	if err := user.HashPassword(body.Password); err != nil {
		ac.logService.Errorf(tags, "Failed to hash password for user %s: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := ac.userService.Create(user); err != nil {
		ac.logService.Errorf(tags, "Failed to create user %s: %v", user.ID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ac.authService.SendWelcomeAndVerif(user)
	ac.logService.Printf(tags, "User %s registered successfully", user.ID)
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Logout godoc
//
//	@Summary		Logout user
//	@Description	Logout user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		201
//	@Failure		400	{object}	utils.HttpError
//	@Router			/api/auth/logout [post]
func (ac *AuthController) Logout(c *gin.Context) {
	token := c.MustGet("token").(*token.Token)
	ac.tokenService.Delete(token.ID)
	ac.authService.ClearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// Refresh godoc
//
//	@Summary		Refresh token
//	@Description	Refresh token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	token.RefreshSuccessResponse
//	@Failure		400		{object}	utils.HttpError
//	@Router			/api/auth/refresh [post]
func (ac *AuthController) Refresh(c *gin.Context) {
	tags := []string{"AuthController", "Refresh"}

	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		ac.logService.Errorf(tags, "Refresh token not found in cookie: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token not found"})
		return
	}

	token, err := ac.tokenService.FindOneBy("refresh_token", refreshTokenString)
	if err != nil || token == nil {
		ac.logService.Errorf(tags, "Refresh token not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	jwtKey := ac.dotenvService.Get("JWT_SECRET_KEY")
	if err := token.GenerateAccessToken(jwtKey); err != nil {
		ac.logService.Errorf(tags, "Failed to generate access token for user ID %s: %v", token.UserID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	if err := ac.tokenService.Update(token); err != nil {
		ac.logService.Errorf(tags, "Failed to update token for user ID %s: %v", token.UserID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
		return
	}

	ac.authService.SetAuthCookies(c, token)
	c.JSON(http.StatusOK, gin.H{"access_token": token.AccessToken})
}
