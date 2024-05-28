package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/torexanovich/savelife/backend/internal/auth"
	"github.com/torexanovich/savelife/backend/internal/database"
)

func main() {
	// Initialize the database
	dbPath := "savelife.db"
	if err := database.Init(dbPath); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize repositories
	userRepo := database.NewUserRepository(database.DB)

	// Initialize authentication service
	jwtKey := []byte("mysupersecretkeythatnoonecanfindyes") // Change this to a more secure key in production
	authService := auth.NewService(userRepo, jwtKey)

	// Initialize authentication middleware
	// authMiddleware := auth.NewMiddleware(authService)

	// Initialize the HTTP router
	router := mux.NewRouter()

	// Register authentication routes
	authHandler := auth.NewAuthHandler(authService)
	router.HandleFunc("/login", authHandler.Login).Methods("POST")

	// // Example protected route
	// router.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("This is a protected route"))
	// }).Methods("GET").Middleware(authMiddleware.Authenticate)

	// Start the HTTP server
	serverAddr := ":8080" // Change this to your desired server address
	http.Handle("/", router)
	log.Printf("Server listening on %s...", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
