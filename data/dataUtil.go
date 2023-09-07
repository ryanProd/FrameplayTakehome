package data

import (
	"fmt"

	"github.com/ryanProd/FrameplayTakehome/structs"
)

// Checks to see if any fields are empty or invalid
func ValidateUsers(users []structs.User) (bool, error) {
	for i, user := range users {
		if user.User_id <= 0 {
			return false, fmt.Errorf("user at index %d had invalid user_id of %d", i, user.User_id)
		}
		if len(user.Username) <= 0 {
			return false, fmt.Errorf("user at index %d has no username field", i)
		}
		if len(user.Email) <= 0 {
			return false, fmt.Errorf("user at index %d has no email field", i)
		}
		if len(user.Password) <= 0 {
			return false, fmt.Errorf("user at index %d has no password field", i)
		}
		if len(user.Created_on) <= 0 {
			return false, fmt.Errorf("user at index %d has no created_on field", i)
		}
	}
	return true, nil
}
