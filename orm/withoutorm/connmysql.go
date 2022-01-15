package withoutorm

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func ConnMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@54321@tcp(127.0.0.1:3306)/ApiDB?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InsertDB(db *sql.DB) (sql.Result, error) {
	// insert data
	//query := "insert into product (name, description, price) values (xx,xx,xx)"
	//result, err := db.Exec(query)
	statement, err := db.Prepare("insert into product (name, description, price) values (?,?,?)")
	result, err := statement.Exec("pen", "used to write something", "2")
	if err != nil {
		return nil, err
	}
	return result, err
}

func QueryDB(db *sql.DB) ([]Product, error) {
	// query data
	query := ""
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	var product []Product
	var name, description, price string
	for rows.Next() {
		if err := rows.Scan(&name, &description, &price); err != nil {
			fmt.Errorf("query err", err.Error())
		}
		product = append(product, Product{
			Name:        name,
			Description: description,
			Price:       price,
		})
	}
	return product, nil
}
