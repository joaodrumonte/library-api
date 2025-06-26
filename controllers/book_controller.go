package controllers

import (
	"context"
	"fmt"
	"library-api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}

func CriarLivro(c *gin.Context) {
	var livro models.Livro

	var raw map[string]interface{}
	if err := c.ShouldBind(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos os campos (titulo, autor, desc) são obrigatórios"})
		return
	}
	if err := c.ShouldBindJSON(&raw); err != nil {

		c.Request.ParseForm()
		raw = make(map[string]interface{})
		for key := range c.Request.PostForm {
			raw[key] = c.Request.PostForm.Get(key)
		}
	}

	camposPermitidos := map[string]bool{"titulo": true, "autor": true, "desc": true}
	for campo := range raw {
		if !camposPermitidos[campo] {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Campo '%s' não permitido", campo)})
			return
		}
	}

	log.Printf("Livro recebido: %+v\n", livro)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, livro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

func ListarLivros(c *gin.Context) {
	var livros []models.Livro
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var livro models.Livro
		if err := cursor.Decode(&livro); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		livros = append(livros, livro)
	}

	if len(livros) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhum livro encontrado. Cadastre um novo."})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func ObterLivro(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'id' é obrigatório"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var livro models.Livro
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&livro)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, livro)
}

func AtualizarLivro(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'id' é obrigatório"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var livroExistente models.Livro
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&livroExistente)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	if err := c.Request.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
		log.Println("Erro ao fazer parse do form:", err)
	}

	titulo := c.PostForm("titulo")
	autor := c.PostForm("autor")
	desc := c.PostForm("desc")

	camposPermitidos := map[string]bool{"titulo": true, "autor": true, "desc": true}
	for key := range c.Request.PostForm {
		if !camposPermitidos[key] {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Campo '%s' não permitido", key)})
			return
		}
	}

	if titulo == "" && autor == "" && desc == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Pelo menos um campo (titulo, autor ou desc) deve ser enviado para atualização",
		})
		return
	}

	if titulo != "" {
		livroExistente.Titulo = titulo
	}
	if autor != "" {
		livroExistente.Autor = autor
	}
	if desc != "" {
		livroExistente.Desc = desc
	}

	update := bson.M{
		"$set": bson.M{
			"titulo": livroExistente.Titulo,
			"autor":  livroExistente.Autor,
			"desc":   livroExistente.Desc,
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado para atualização"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro atualizado com sucesso"})
}

func DeletarLivro(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'id' é obrigatório"})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado para exclusão"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Livro deletado com sucesso"})
}
