package main // import "github.com/cloudspace/microservice-postgres"

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	db_host, db_port, db_user, db_password, db_name, query := os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5], os.Args[6]

	sqlOpenRequest := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db_host, db_port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", sqlOpenRequest)
	if err != nil {
		log.Fatal(err)
		return
	}

	result, err := getJSON(query, db)
	if err != nil {
		fmt.Println(getJsonError(err))
	}
	db.Close()

	var wrappedResults = map[string]string{
		"result": result,
	}

	encodedResults, _ := json.Marshal(wrappedResults)
	fmt.Println(string(encodedResults))
}

func getJSON(sqlString string, db *sql.DB) (string, error) {
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func getJsonError(myError error) string {

	errorJson := make(map[string]interface{})
	errorJson["error"] = myError.Error()
	jsonData, err := json.Marshal(errorJson)
	if err != nil {
		return "{\"error\": \"There was an error generatoring the error.. goodluck\"}"
	}
	return string(jsonData)
}
