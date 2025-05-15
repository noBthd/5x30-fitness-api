package services

import (
	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/noBthd/5x30-fitness-api/internal/models"
)

func GetAllUsers()([]models.User, error) {

    rows, err := db.DB.Query("SELECT id, username, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []models.User
    for rows.Next() {
        var user models.User
        err := rows.Scan(&user.ID, &user.Username, &user.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}
