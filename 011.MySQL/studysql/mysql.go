package main

import (
	"database/sql"
	"fmt"

	"github.com/ani8570/learngo/010.file/f"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:dlwlgns2682!@tcp(127.0.0.1:3306)/pets")
	f.CheckErr(err)
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM cats WHERE owner = 'Casey'").Scan(&name)
	f.CheckErr(err)
	rows, err = db.Query("SELECT id, name FROM cas", 1)
	f.CheckErr(err)
	defer rows.Close()

	var id int

	for rows.Next() {
		err := rows.Scan(&d, &name)
		
	}
	mt.Println(name)
}
