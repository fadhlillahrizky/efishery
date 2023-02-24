package main

import (
	"efishery-golang/middleware"
	"efishery-golang/modules/auth"
	"efishery-golang/modules/resources"
	"efishery-golang/utils/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {

	router := gin.Default()
	router.Use(
		middleware.AllowCORS(),
	)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	apiRouter := router.Group("/api")
	config := os.Getenv("DB_MYSQL_DSN")
	DB := db.GormMysql(config)

	authRouter := apiRouter.Group("/auth")
	authHandler := auth.NewRequestHandler(DB)
	authHandler.Handle(authRouter)

	resourcesRouter := apiRouter.Group("/resources")
	resourcesHandler := resources.NewRequestHandler(DB)
	resourcesHandler.Handle(resourcesRouter)

	err := router.Run()

	if err != nil {
		log.Println("main router.Run:", err)
		return
	}
}
