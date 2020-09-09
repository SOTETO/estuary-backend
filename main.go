package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/tombiers/estuary-backend/router"
)

type AnotherTable struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

func main() {
	router.SayHi()
	router.HandleRequests()
}
