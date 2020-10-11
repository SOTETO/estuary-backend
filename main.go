package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/tombiers/estuary-backend/router"
)

func main() {
	router.HandleRequests()
}
