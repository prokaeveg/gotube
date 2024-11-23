package db

import (
	"database/sql"
	"fmt"
	"net/http"
)

func ShowTables(db *sql.DB, w http.ResponseWriter) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		http.Error(w, "Failed to execute query", http.StatusInternalServerError)
		//@todo logging
		return
	}

	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			//@todo logging
			continue
		}
		tables = append(tables, table)
	}

	if len(tables) == 0 {
		w.Write([]byte("No tables found"))
	}

	w.Write([]byte(fmt.Sprintf("Tables: %v", tables)))
}
