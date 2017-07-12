package mmain
import (
         "database/sql"
         _ "github.com/go-sql-driver/mysql"
         "fmt"
         )

func main(){
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		fmt.Println("not connected to DB")
	}
	fmt.Println("connected to DB")
	defer db.Close()

	err = db.Ping()
	if err != nil{
		fmt.Println("Error", err)
	}
	stmtOut, err := db.Prepare("SELECT division FROM mactovendor WHERE macAddress =?")
	if err != nil{
		panic(err.Error())
	}
	defer stmtOut.Close()
var div string
err = stmtOut.QueryRow(1).Scan(&div)
if err != nil{
		panic(err.Error())
	}
	fmt.Println("name of the division is %d" , div)
}

