package operations

import (
	"CRUD_API"
	"CRUD_API/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Create(c *gin.Context) {
	var products CRUD_API.Products
	if err := c.BindJSON(&products); err != nil { //привязка данных из запроса к нашей структуре
		c.JSON(http.StatusBadRequest, "err in create ") //ошибка в формате json
		return
	}
	result := database.DB.Create(&products)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func ReadAll(c *gin.Context) {
	var products []CRUD_API.Products //срез чтобы хранить не 1,а несколько
	err := database.DB.Find(&products)
	if err != nil {
		log.Fatalf("err in read:%s", err.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})

}
func Read(c *gin.Context) {
	id := c.Param("id")
	var products CRUD_API.Products
	result := database.DB.First(&products, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
func Update(c *gin.Context) {

}
func Delete(c *gin.Context) {

}
