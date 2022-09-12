package authors

import "github.com/gin-gonic/gin"

func AuthorsRegister(router *gin.RouterGroup) {
	router.GET("authors/:id", AuthorGet)
	router.POST("authors", AuthorCreate)
	router.DELETE("authors/:id", AuthorDelete)
	router.PUT("authors/:id", AuthorUpdate)
	router.GET("authors", AuthorsFind)
}

func AuthorGet(c *gin.Context) {
	
}

func AuthorUpdate(c *gin.Context) {

}

func AuthorCreate(c *gin.Context) {

}

func AuthorDelete(c *gin.Context) {

}

func AuthorsFind(c *gin.Context){

}
