package models

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Tool struct {
	Id int         `json:"Id"`
	Title string   `json:"Title"`
	Link string    `json:"Link"`
	Descrip string `json:"Description"` 
	Tags []string  `json:"tags"`
}

type Tag struct {
	Id int        `json:"id`
	Tags []string  `json:"tags"`
}

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

func (h *BaseHandler) InsertTool(T Tool) Tool{

	stmt, err := h.db.Prepare("INSERT INTO tool (title, link, descrip) VALUES (?,?,?)")
	if err != nil {
		log.Println("failed to prepare statement")
	}
	defer stmt.Close()

	res, err := stmt.Exec(T.Title, T.Link, T.Descrip)
	if err != nil {
		log.Println(err)
	}
	Id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return Tool{}
	}
	InsertTags(h, Id, T.Tags)

	T.Id = int(Id)
	return T
}

func InsertTags(h *BaseHandler, Id int64, Tags []string) {
	sqlSTR := "INSERT INTO" + "`key`" + "(id_tool, tag) VALUES"
	values := []interface{}{}

	for _, row := range Tags{
		sqlSTR += "(?, ?),"
		values = append(values, Id, row)
	}

	sqlSTR = sqlSTR[0:len(sqlSTR)-1]
	stmt, err := h.db.Prepare(sqlSTR)
	if err != nil {
		log.Println(err)
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		log.Println(err)
	}
}

func (h *BaseHandler) GetTools() []Tool {
	rows, err := h.db.Query("SELECT tool.id_tools, tool.title, tool.link, tool.descrip," + "`key`.`tag`" + "FROM tool INNER JOIN" + "`key`" + "ON tool.id_tools = key.id_tool")
	if err != nil {
		log.Println(err)
	}
	var T []Tool
	
	var Qtd int = 0

	for rows.Next(){
		temp := struct{Id int
			 Title string
			 Link string
			 Descrip string
			 Tag string}{}
		err := rows.Scan(&temp.Id, &temp.Title, &temp.Link, &temp.Descrip, &temp.Tag)
		
		if len(T) == 0 || T[len(T)-1].Id != temp.Id{
			T = append(T, Tool{temp.Id, temp.Title, temp.Link, temp.Descrip, []string{temp.Tag}})
		} else {
			T[len(T)-1].Tags = append(T[len(T)-1].Tags, temp.Tag)
		}
		Qtd++

		if err != nil {
			log.Println(err)
		}
	}
	return T
}

func (h *BaseHandler) DeleteTool(Id int) bool {
	stmt, err := h.db.Prepare("DELETE FROM main.tool WHERE id_tools = (?)")
	if err != nil {
		log.Println(err)
	}

	res, err := stmt.Exec(Id)
	if err != nil {
		log.Println(err)
	}

	RA, _ := res.RowsAffected()
	if RA != 0 {
		return true
	} 
	return false
}