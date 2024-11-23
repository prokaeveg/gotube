package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

func ShowTables(db *pgxpool.Pool, w http.ResponseWriter) {
	rows, err := db.Query(context.Background(), "SELECT tablename FROM pg_catalog.pg_tables")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		return
	}

	w.Write([]byte(fmt.Sprintf("Tables: %v", tables)))
}
