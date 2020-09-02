package main


import(
	"net/http"
	"gin-crud/models"
	"gin-crud/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})

	})

	models.ConnectDateBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:title", controllers.FindBook)
	r.PATCH("/books/:title", controllers.UpdateBook)
	r.DELETE("/books/:title", controllers.DeleteBook)

	r.Run()
}