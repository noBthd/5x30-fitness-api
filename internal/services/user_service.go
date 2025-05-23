package services

import (
	"errors"

	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/noBthd/5x30-fitness-api/internal/models"
)

func GetAllUsers()([]models.User, error) {
    rows, err := db.DB.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.RegDate, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func UserExists(email string)([]models.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users WHERE users.email = '" + email + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.RegDate, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
	}

	return users, nil
}

func CreateUser(user models.User)(error) {
	if len(user.Email) == 0 {
		return errors.New("email cannot be empty")
	}

	_, err := db.DB.Query("INSERT INTO users (email, password, registration_date) VALUES ($1, $2, NOW())", 
		user.Email, string(user.Hashed_password))
	if err != nil {
		return err
	}

	return nil
}