package code

import (
	"fmt"
	"golang-api/core"
	ginM "golang-api/gin"
	"golang-api/log"
	"golang-api/middleware"
	"golang-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ICodeController interface {
	core.IProvider
	RegisterRoutes()
	Verify(c *gin.Context)
	RequestCode(c *gin.Context)
	ResetPassword(c *gin.Context)
	RequestPasswordReset(c *gin.Context)
}

type CodeController struct {
	*core.Provider
	codeService ICodeService
	userService user.IUserService
	ginService  ginM.IGinService
	logService  log.ILogService
}

func NewCodeController(module *CodeModule) *CodeController {
	return &CodeController{
		Provider:    core.NewProvider("CodeController"),
		codeService: module.Get("CodeService").(ICodeService),
		userService: module.Get("UserService").(user.IUserService),
		ginService:  module.Get("GinService").(ginM.IGinService),
		logService:  module.Get("LogService").(log.ILogService),
	}
}

func (cc *CodeController) OnInit() error {
	cc.RegisterRoutes()
	return nil
}

func (cc *CodeController) RegisterRoutes() {
	fmt.Println("Registering Code routes")
	code := cc.ginService.GetGroup().Group("/code")
	code.POST("/verify",
		middleware.Validate[VerifyUserDto](),
		cc.Verify,
	)
	code.POST("/request-verify",
		middleware.Validate[RequestCodeDto](),
		cc.RequestCode,
	)
	code.POST("/reset",
		middleware.Validate[ResetPasswordDto](),
		cc.ResetPassword,
	)
	code.POST("/request-reset",
		middleware.Validate[RequestCodeDto](),
		cc.RequestPasswordReset,
	)
}

// Verify godoc
//
//	@Summary		Verify account
//	@Description	Verify user account with code
//	@Tags			code
//	@Accept			json
//	@Produce		json
//	@Param			code	body		code.VerifyUserDto	true	"Verify user"
//	@Success		200
//	@Failure		400	{object}	utils.HttpError
//	@Router			/api/code/verify [post]
func (cc *CodeController) Verify(c *gin.Context) {
	tags := []string{"CodeController", "Verify"}
	body, _ := c.MustGet("body").(VerifyUserDto)

	code, err := cc.codeService.FindOneByCodeAndEmail(body.Code, body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "Code not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Code not found"})
		return
	}

	if code.IsExpired() {
		cc.logService.Errorf(tags, "Code expired for email: %s", body.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired code"})
		return
	}

	user, err := cc.userService.FindOneBy("email", body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "User %s not found: %v", body.Email, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Verified = true

	if err := cc.userService.Update(user); err != nil {
		cc.logService.Errorf(tags, "Failed to verify account for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify account"})
		return
	}

	if err := cc.codeService.Delete(code.ID); err != nil {
		cc.logService.Errorf(tags, "Failed to delete code for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete code"})
		return
	}

	cc.logService.Printf(tags, "User %s verified successfully", user.Email)
	c.JSON(http.StatusOK, gin.H{"message": "Account verified successfully"})
}

// RequestCode godoc
//
//	@Summary		Request code
//	@Description	Request code
//	@Tags			code
//	@Accept			json
//	@Produce		json
//	@Param			request	body	code.RequestCodeDto	true	"Request code"
//	@Success		200
//	@Failure		400	{object}	utils.HttpError
//	@Router			/api/code/request-verify [post]
func (cc *CodeController) RequestCode(c *gin.Context) {
	body, _ := c.MustGet("body").(RequestCodeDto)
	tags := []string{"CodeController", "RequestCode"}

	user, err := cc.userService.FindOneBy("email", body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "Email %s not found: %v", body.Email, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	if user.Verified {
		cc.logService.Errorf(tags, "Account already verified for user %s", user.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Account already verified"})
		return
	}

	code := NewCode(user.ID, user.Email)
	code.GenerateCode()

	if err := cc.codeService.DeleteBy("user_id", user.ID); err != nil {
		cc.logService.Errorf(tags, "Failed to delete existing codes for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing codes"})
		return
	}

	if err := cc.codeService.Create(code); err != nil {
		cc.logService.Errorf(tags, "Failed to create code for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create code"})
		return
	}

	if err := cc.codeService.SendVerifCodeEmail(user.Email, code.Code); err != nil {
		cc.logService.Errorf(tags, "Failed to send code email to user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send code email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Code sent successfully"})
}

// ResetPassword godoc
//
//	@Summary		Reset password
//	@Description	Reset user password
//	@Tags			code
//	@Accept			json
//	@Produce		json
//	@Param			reset	body		code.ResetPasswordDto	true	"Reset password"
//	@Success		200
//	@Failure		400	{object}	utils.HttpError
//	@Router			/api/code/reset [post]
func (cc *CodeController) ResetPassword(c *gin.Context) {
	body, _ := c.MustGet("body").(ResetPasswordDto)
	tags := []string{"CodeController", "ResetPassword"}

	code, err := cc.codeService.FindOneByCodeAndEmail(body.Code, body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "Verification code not found: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Verification code not found"})
		return
	}

	if code.IsExpired() {
		cc.logService.Errorf(tags, "Verification code expired for email: %s", body.Email)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or expired verification code"})
		return
	}

	user, err := cc.userService.FindOneBy("email", body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "User %s not found: %v", body.Email, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.HashPassword(body.Password) != nil {
		cc.logService.Errorf(tags, "Failed to hash password for user %s", user.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	if err := cc.userService.Update(user); err != nil {
		cc.logService.Errorf(tags, "Failed to reset password for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	if err := cc.codeService.Delete(code.ID); err != nil {
		cc.logService.Errorf(tags, "Failed to delete verification code for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete verification code"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// RequestPasswordReset godoc
//
//	@Summary		Request password reset
//	@Description	Request password reset
//	@Tags			code
//	@Accept			json
//	@Produce		json
//	@Param			request	body	code.RequestCodeDto	true	"Request password reset"
//	@Success		200
//	@Failure		400	{object}	utils.HttpError
//	@Router			/api/code/request-reset [post]
func (cc *CodeController) RequestPasswordReset(c *gin.Context) {
	body, _ := c.MustGet("body").(RequestCodeDto)
	tags := []string{"CodeController", "RequestPasswordReset"}

	user, err := cc.userService.FindOneBy("email", body.Email)
	if err != nil {
		cc.logService.Errorf(tags, "Email %s not found: %v", body.Email, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
		return
	}

	code := NewCode(user.ID, user.Email)
	code.GenerateCode()

	if err := cc.codeService.DeleteBy("user_id", user.ID); err != nil {
		cc.logService.Errorf(tags, "Failed to delete existing reset codes for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete existing reset codes"})
		return
	}

	if err := cc.codeService.Create(code); err != nil {
		cc.logService.Errorf(tags, "Failed to create password reset code for user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create password reset code"})
		return
	}

	if err := cc.codeService.SendResetCodeEmail(user.Email, code.Code); err != nil {
		cc.logService.Errorf(tags, "Failed to send code email to user %s: %v", user.Email, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send code email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset code sent successfully"})
}
