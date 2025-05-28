package services

import (
	"errors"

	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/noBthd/5x30-fitness-api/internal/models"
	"golang.org/x/crypto/bcrypt"
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
        err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RegDate)
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
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RegDate)
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

func SignIn(userEP models.User)(bool, []models.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users WHERE users.email = $1",
		userEP.Email)
	if err != nil {
		return false, nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RegDate)
        if err != nil {
            return false, nil, err
        }

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userEP.Password))
		if err != nil {
			return false, nil, err
		}

        users = append(users, user)
	}


	return true, users, nil
}

func GetUser(email string)([]models.User, error) {
	rows, err := db.DB.Query("SELECT * FROM users WHERE email = $1",
		email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.RegDate)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	
	return users, nil
}