package main

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CustomType struct {
	V string
}

func (c CustomType) Value() (driver.Value, error) {
	//return c.V, nil
	return "fake", errors.New("fake error")
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(db_url:3306)/db_name")
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
		return
	}
	defer db.Close()

	customValue := CustomType{"123"}
	rows, err := db.Query("SELECT bundle_id FROM `table_name` WHERE user_id = ? limit 10", customValue)
	if err != nil {
		fmt.Printf("Error querying database: %v", err)
		return
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return
		}
		fmt.Println("Result:", result)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterating rows: %v", err)
		return
	}

	fmt.Println("Done")
}
