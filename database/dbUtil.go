package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ryanProd/FrameplayTakehome/config"
	"github.com/ryanProd/FrameplayTakehome/structs"
)

// Establishes conenction to database. The connection string is formed from parsing the .env variables
func ConnectDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"), config.Config("DB_HOST"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}

	return db
}

/*
With an array of unique user_id's, queries the database and
returns the corresponding user data as an array of User structs
*/
func QueryDBforUsers(db *sql.DB, ids []int) ([]structs.User, error) {
	var users []structs.User

	for _, id := range ids {
		var user structs.User
		if err := db.QueryRow("SELECT * FROM accounts WHERE user_id = $1", id).Scan(&user.User_id,
			&user.Username, &user.Password, &user.Email, &user.Created_on); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil

}
