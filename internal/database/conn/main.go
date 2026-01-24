package dbConn

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	database "github.com/attendeee/event-app/internal/database/compiled-sql"
)

var db *sql.DB
var Query *database.Queries
var Context context.Context

func init( /* Skibidi */ ) {
	var err error

	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic(fmt.Sprintf("%s %s", "Unable to open database connection", err))
	}

	Query = database.New(db)

	Context = context.Background()

}
