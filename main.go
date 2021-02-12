package main

import (
	"fmt"
	"github.com/george124816/api-vuttr/internal/pkg/models"
	_ "github.com/go-sql-driver/mysql"

	
)

func main(){
	
	tools := models.Tool{
		1, 
		"Notion", 
		"https://notion.so", 
		"All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized. ",
		[]string{"organization", "planing"},

	}
	fmt.Println(tools)
}



