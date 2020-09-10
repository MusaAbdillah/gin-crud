package main


import(
	"net/http"
	"gin-crud/models"
	"gin-crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	
	_ "gin-crud/docs" // docs is generated by Swag CLI, you have to import it.

)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main(){
	r := gin.Default()


	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})

	})

	models.ConnectDateBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)
	r.POST("/users/register", controllers.Register)
	r.GET("/users", controllers.Index)
	r.GET("/users/:id", controllers.Show)

	// swagger url
	// url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run()
}