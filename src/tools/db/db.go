package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/cn-lxy/music-api/tools/config"
	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB

func init() {
	dbInit()
	log.Println("Dabase connection pool initialized")
	log.Println("check db server")
	err := checkDbServer()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Initialize the database connection pool.
func dbInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		config.Cfg.Db.UserName,
		config.Cfg.Db.Password,
		config.Cfg.Db.Host,
		config.Cfg.Db.Port,
		config.Cfg.Db.Name)
	log.Println(dsn)
	var err error
	pool, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)
}

// Check Database whether is open.
func checkDbServer() error {
	sql := "SELECT VERSION();"
	res, err := Query(sql)
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(res) == 0 {
		return fmt.Errorf("mysql DataBase Server not start")
	}
	return nil
}

// Prepare statement for reading data
func Query(sqlString string, args ...any) ([]map[string]any, error) {
	// Execute the query
	rows, err := pool.Query(sqlString, args...)
	if err != nil {
		// proper error handling instead of panic in your app
		return nil, err
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		// proper error handling instead of panic in your app
		return nil, err
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	dataSlice := make([]map[string]any, 0)

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			// proper error handling instead of panic in your app
			return nil, err
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		oneData := make(map[string]any, len(columns))
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			oneData[columns[i]] = value
			// fmt.Println(columns[i], ": ", value)
		}
		dataSlice = append(dataSlice, oneData)
		// fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		// proper error handling instead of panic in your app
		return nil, err
	}
	return dataSlice, nil
}

// Update database, the `update`, `insert`, and `delete` both is this function.
func Update(sqlString string, params ...any) (int64, error) {
	result, err := pool.Exec(sqlString, params...)
	if err != nil {
		return 0, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, fmt.Errorf("%s", "update error")
	}
	return result.LastInsertId()
}
