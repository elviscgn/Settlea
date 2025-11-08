package main

import (
	router "settlea/api"
	"settlea/internal/ws"
)

// "log"

// "settlea/internal/ws"

func main() {
	// dbConn, err := db.NewDatabase()
	// if err != nil {
	// log.Fatalf("could not initialize database connection: %s", err)
	// }

	// userRep := user.NewRepository(dbConn.GetDB())
	// userSvc := user.NewService(userRep)
	// userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	// router.InitRouter()
	router.InitRouter(wsHandler)
	router.Start("localhost:8080")
}
