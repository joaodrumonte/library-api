package routes

import (
	"library-api/controllers"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	router.POST("/livros", controllers.CriarLivro)
	router.GET("/livros", controllers.ListarLivros)
	router.GET("/livro", controllers.ObterLivro)
	router.PUT("/livro", controllers.AtualizarLivro)
	router.DELETE("/livro", controllers.DeletarLivro)
}
