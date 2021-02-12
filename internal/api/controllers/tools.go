package controllers

import (
	"github.com/george124816/api-vuttr/internal/pkg/db"
	"github.com/george124816/api-vuttr/internal/pkg/models"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/george124816/my-own-migration/migrate"
	"log"
	"strconv"
)

var database *sql.DB
var handle *models.BaseHandler

func init(){
	database = db.ConnectDB()
	handle = models.NewBaseHandler(database)
	migrate.Migrate(database, 2)

}

func Hello(c *gin.Context){
	c.JSON(200, "Hello")
}

func InsertTool(c *gin.Context){
	var T models.Tool
	c.Bind(&T)
	log.Println(T)
	if T.Title != "" && T.Link != "" && T.Descrip != "" {
		T = handle.InsertTool(T)
		c.JSON(201, T)
		return
	}
	c.JSON(400, gin.H{"error": "Fields are empty"})
}

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