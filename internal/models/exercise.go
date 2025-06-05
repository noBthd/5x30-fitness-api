package models

type Exercise struct {
	ID 		int		`json:"id"`
	Name 	string	`json:"name"`
	Dur 	string	`json:"duration"`
	Reps 	string	`json:"reps"`
	Rest 	string	`json:"rest"`
}