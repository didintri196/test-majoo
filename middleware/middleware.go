package middleware

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	cors "github.com/itsjamie/gin-cors"

	"majoo-test/config"

	"majoo-test/controllers"

	_ "majoo-test/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var (
	productcontroll = new(controllers.ProductController)
	usercontroll    = new(controllers.UserController)
)

func Middleware() {

	flag.Parse()
	dir := config.GetDir()

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))
	router.Static("/public", dir+"/../assets/")

	linkproduct := router.Group("/v1/product")
	{
		linkproduct.GET("", auth, productcontroll.ListHandler)
		linkproduct.GET("/view/:id", auth, productcontroll.ShowHandler)
		linkproduct.POST("", auth, productcontroll.CreateHandler)
		linkproduct.PUT("/:id", auth, productcontroll.PutHandler)
		linkproduct.DELETE("/:id", auth, productcontroll.DeleteHandler)
	}

	linkuser := router.Group("/v1/user")
	{
		linkuser.GET("/session", auth, usercontroll.ShowSessionHandler)
		linkuser.POST("/login", usercontroll.UserLogin)

		linkuser.POST("", auth, usercontroll.CreateHandler)
		linkuser.GET("", auth, usercontroll.ListHandlerMercant)
		linkuser.GET("/view/:id", auth, usercontroll.ShowHandler)
		linkuser.PUT("/:id", auth, usercontroll.PutHandler)
		linkuser.DELETE("/:id", auth, usercontroll.DeleteHandler)
	}

	urlswag := ginSwagger.URL("doc.json")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, urlswag))

	router.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "API MINI-POST ON"})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "NOT FOUND"})
	})
	router.Run(":3333")

}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if token != nil && err == nil {
		claims := jwt.MapClaims{}
		_, _ = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

	} else {
		result := gin.H{
			"error":   err.Error(),
			"message": "Tolong Masukan Authorization",
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}

}
