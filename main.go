package main

import (
	"github.com/server/auth/infra/database"
	"github.com/server/auth/infra/handlers"
)

func main() {
	handlers.NewAuthHandlers(database.NewConnectionDB()).Start(3001)
}
