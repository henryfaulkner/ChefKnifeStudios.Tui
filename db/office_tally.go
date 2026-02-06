package db

import (
	"database/sql"
	"time"
)

type OfficeTally struct {
	IID         int64
	SRestaurant string
	ITally      int64
	DtCreatedOn time.Time
	BIsDeleted  bool
}

func InsertOfficeTally(db *sql.DB, tally *OfficeTally) (int64, error) {
	createdOn := time.Now().UTC().Format(time.RFC3339)
	result, err := db.Exec(
		"INSERT INTO officetally (srestaurant, itally, dtcreatedon, bisdeleted) VALUES (?, ?, ?, ?)",
		tally.SRestaurant, tally.ITally, createdOn, boolToInt(tally.BIsDeleted),
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
