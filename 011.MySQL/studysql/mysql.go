package main

import (
	"database/sql"
	"fmt"

	"github.com/ani8570/learngo/010.file/f"
	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	id   int
	Name string
}

func printIdandName(rows *sql.Rows) {
	var name string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &name)
		f.CheckErr(err)
		fmt.Println(id, name)
	}
}

func insertdata(db *sql.DB) {
	c := person{4, "Judack"}
	result, err := db.Exec("INSERT IGNORE INTO test VALUES (?,?)", c.id, c.Name) //, c.Owner, c.birthDay)
	f.CheckErr(err)

	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("One row inserted")
	}
}
func main() {
	db, err := sql.Open("mysql", "root:dlwlgns2682!@tcp(127.0.0.1:3306)/testdb")
	f.CheckErr(err)
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM test WHERE id >= ?", 2).Scan(&name)
	f.CheckErr(err)

	fmt.Println(name)

	rows, err := db.Query("SELECT id, name FROM test")
	f.CheckErr(err)
	defer rows.Close()
	printIdandName(rows)

	insertdata(db)
	rows, err = db.Query("SELECT id, name FROM test")
	f.CheckErr(err)
	defer rows.Close()

	printIdandName(rows)

	stmt, err := db.Prepare("UPDATE test SET name=? WHERE id=?")
	f.CheckErr(err)
	defer stmt.Close()

	_, err = stmt.Exec("Tom", 1)
	f.CheckErr(err)
	_, err = stmt.Exec("Jack", 2)
	f.CheckErr(err)
	_, err = stmt.Exec("SHawn", 3)
	f.CheckErr(err)

}
