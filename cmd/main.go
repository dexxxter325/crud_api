package main

import (
	"CRUD_API/database"
	"CRUD_API/operations"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "World",
		})
	})
	database.ConnectToDb()
	/*err := database.DB.AutoMigrate(&CRUD_API.Product{}) //автоматически создает миграции
	if err != nil {
		log.Fatalf("failed in automigrate %s", err.Error())
	}*/
	router.POST("/product", operations.Create)
	router.GET("/product", operations.ReadAll)
	router.GET("/product/:id", operations.Read)
	router.PUT("/product", operations.Update)
	router.DELETE("/product", operations.Delete)
	errr := router.Run(":8080")
	if errr != nil {
		log.Fatalf("error in router.run %s", errr.Error())
	}

}
