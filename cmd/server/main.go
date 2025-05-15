package main

import (
	"github.com/noBthd/5x30-fitness-api/internal/db"
	"github.com/yourname/5x30-fitness-api/internal/db"
)

func main() {
	cfg := cfg.GetConfig()
	db.DBInit(cfg)
	
}
