package controllers

import (
	"github.com/george124816/api-vuttr/internal/pkg/db"
	"github.com/george124816/api-vuttr/internal/pkg/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var database *sql.DB
var handle *models.BaseHandler

func init(){
	database = db.ConnectDB()
	Migrate(database)
	handle = models.NewBaseHandler(database)
}

// InsertTool godoc
// @Summary Cadastra Ferramenta
// @Description Cadastra uma nova ferramenta
// @Accept json
// @Param user body models.Tool true "Tool"
// @Success 201 {object} models.Tool
// @Error 400
// @Router /tool/ [post]
func InsertTool(c *gin.Context){
	var T models.Tool
	c.Bind(&T)
	if T.Title != "" && T.Link != "" && T.Descrip != "" {
		T = handle.InsertTool(T)
		c.JSON(201, T)
		return
	}
	c.JSON(400, gin.H{"error": "Fields are empty"})
}

// GetTools godoc
// @Summary Lista Ferramentas
// @Description Lista todas as ferramentas ou ferramentas pela tag
// Produce json
// @Param tag query string false "search tool by tag"
// @Success 200 200 {array} models.Tool
// @Router /tools [get]
func GetTools(c *gin.Context){
	T := handle.GetTools()
	var TR []models.Tool
	
	if c.Query("tag") != "" {
		tag := c.Query("tag")
		for i, j:= range T{
			for _, v := range j.Tags {
				if v == tag {
					TR = append(TR, T[i])
				}
			} 
		}
		c.JSON(200, TR)
	} else {
		c.JSON(200, T)
	}
}

// DeleteTool godoc
// @Summary Deleta Ferramenta
// @Description Deleta uma Ferramenta pelo Id
// @Accept json
// @Param id path int true "Tool ID"
// @Success 200 
// @Router /tools/{id} [delete]
func DeleteTool(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(404, "Not Found")
	}
	if handle.DeleteTool(id) == true {
		c.JSON(200, "")
		return
	}
	c.JSON(404, "Element doesn't exist")
}