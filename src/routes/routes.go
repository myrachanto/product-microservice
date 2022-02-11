package routes

import (
	"fmt"

	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/myrachanto/microservice/product/src/controllers"
	"github.com/myrachanto/microservice/product/src/repository"
	service "github.com/myrachanto/microservice/product/src/services"

	"github.com/spf13/viper"
)

//StoreAPI =>entry point to routes
type Open struct {
	Port     string `mapstructure:"PORT"`
	Key      string `mapstructure:"EncryptionKey"`
	DURATION string `mapstructure:"DURATION"`
}

// func LoadConfig(path string) (open Open, err error) {
func LoadConfig() (open Open, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&open)
	return
}
func StoreApi() {
	open, err := LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	repository.IndexRepo.InitDB()
	//check db connection//////////////////////
	fmt.Println("initialization----------------")
	controllers.NewproductController(service.NewproductService(repository.NewproductRepo()))
	e := echo.New()

	e.Static("/", "public")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	JWTgroup := e.Group("/api/")
	JWTgroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(open.Key),
	}))
	e.POST("/create", controllers.ProductController.Create)
	e.GET("/product/:id", controllers.ProductController.GetOne)
	e.GET("/products/:code", controllers.ProductController.GetOnebyCode)
	e.GET("/products", controllers.ProductController.GetAll)
	JWTgroup.PUT("product/:id", controllers.ProductController.Update)
	JWTgroup.DELETE("product/:id", controllers.ProductController.Delete)
	//e.DELETE("loggoutall/:id", controllers.ProductController.DeleteALL) logout all accounts

	// Start server
	e.Logger.Fatal(e.Start(open.Port))
}
