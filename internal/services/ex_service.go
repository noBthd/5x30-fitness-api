package services

import (
	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/noBthd/5x30-fitness-api/internal/models"
)

func GetAllExercises()([]models.Exercise, error) {
	rows, err := db.DB.Query("SELECT exercises_id, exercises_name, ex_dur, ex_reps, ex_rest  FROM exercises")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var exs []models.Exercise
    for rows.Next() {
        var ex models.Exercise
        err := rows.Scan(&ex.ID, &ex.Name, &ex.Dur, &ex.Reps, &ex.Rest)
        if err != nil {
            return nil, err
        }
        exs = append(exs, ex)
    }

    return exs, nil
}