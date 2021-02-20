package studysql

import (
	"database/sql"
	"fmt"

	"github.com/ani8570/learngo/010.file/f"
)

func main() {
	db, err := sql.Open("mysql", "root:pwd@tcp(127.0.0.1:3306)/testdb")
	f.CheckErr(err)
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM test1 WHERE id = 1").Scan(&name)
	f.CheckErr(err)
	fmt.Println(name)

}
