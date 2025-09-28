package gin

import (
	"fmt"
	"golang-api/core"
	"golang-api/docs"
	"golang-api/dotenv"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type IGinService interface {
	core.IProvider
	InitEngine()
	Run()
	Swagger()
	Cors() gin.HandlerFunc
	GetGroup() *gin.RouterGroup
}

type GinService struct {
	*core.Provider
	dotenvService dotenv.IDotenvService
	r             *gin.Engine
	Group         *gin.RouterGroup
}

func NewGinService(module core.IModule) *GinService {
	return &GinService{
		Provider:      core.NewProvider("GinService"),
		dotenvService: module.Get("DotenvService").(dotenv.IDotenvService),
	}
}

func (gs *GinService) OnInit() error {
	gs.InitEngine()
	return nil
}

func (gs *GinService) GetGroup() *gin.RouterGroup {
	return gs.Group
}

func (gs *GinService) Cors() gin.HandlerFunc {
	allowedOrigins := gs.dotenvService.Get("ALLOWED_ORIGINS")

	config := cors.DefaultConfig()
	config.AllowOrigins = strings.Split(allowedOrigins, ",")
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"}
	config.AllowCredentials = true
	config.AllowWebSockets = true
	config.AllowWildcard = true
	config.MaxAge = 0

	return cors.New(config)
}

func (gs *GinService) Swagger() {
	name := gs.dotenvService.Get("NAME")
	doamin := gs.dotenvService.Get("DOMAIN")
	docs.SwaggerInfo.Title = name
	docs.SwaggerInfo.Description = "This is a sample server for " + name
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = doamin
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func (gs *GinService) InitEngine() {
	fmt.Println("Create gin engine")
	ginMode := gs.dotenvService.Get("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode
	}
	gin.SetMode(ginMode)

	gs.r = gin.Default()
	gs.r.Use(gin.Logger())
	gs.r.Use(gin.Recovery())
	gs.r.Use(gs.Cors())

	gs.Group = gs.r.Group("/api")
	gs.Group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (gs *GinService) Run() {
	port := gs.dotenvService.Get("PORT")
	fmt.Printf("Starting server on port %s\n", port)
	if err := gs.r.Run(port); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
