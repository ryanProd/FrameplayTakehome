package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/ryanProd/FrameplayTakehome/config"
	"github.com/ryanProd/FrameplayTakehome/structs"
)

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

	fmt.Printf("version=%s\n", version)

	return db
}

/*
Ideally would add more verification such as handling cases where the user_id is not present in the DB
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
