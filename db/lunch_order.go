package db

import (
	"database/sql"
	"time"
)

type LunchOrder struct {
	IID         int64
	SRestaurant string
	JMetadata   string
	DtCreatedOn time.Time
	BIsDeleted  bool
}

func InsertLunchOrder(db *sql.DB, order *LunchOrder) (int64, error) {
	createdOn := time.Now().UTC().Format(time.RFC3339)
	result, err := db.Exec(
		"INSERT INTO lunchorder (srestaurant, jmetadata, dtcreatedon, bisdeleted) VALUES (?, ?, ?, ?)",
		order.SRestaurant, order.JMetadata, createdOn, boolToInt(order.BIsDeleted),
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
