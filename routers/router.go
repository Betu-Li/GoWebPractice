package routers

import (
	"GinExample2/pkg/logger"
	"GinExample2/routers/controllers"
	"GinExample2/setting"
	"fmt"
	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	if setting.Conf.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.Static("/static", "static")

	router.Use(gin.LoggerWithConfig(logger.LoggerToFile()))
	router.Use(logger.Recover)
	store, err := sessions_redis.NewStore(10, "tcp", setting.RedisAddress, "", []byte("secret"))
	if err != nil {
		fmt.Printf("Failed to create Redis store: %v\n", err)
		panic(err)
	}
	router.Use(sessions.Sessions("mysession", store))

	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.PUT("/login", controllers.UserController{}.Login)
	}

	return router
}
