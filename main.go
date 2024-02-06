package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/swaroop-giri/GoAgg/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	// Load Environment Variables
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to Database: ", err)
	}

	db := database.New(conn)
	apiCfg := &apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	routerv1 := chi.NewRouter()
	routerv1.Get("/health", handlerReadiness)
	routerv1.Get("/err", handlererr)
	routerv1.Post("/users", apiCfg.handlerCreateUser)
	routerv1.Get("/users", apiCfg.authMiddleware(apiCfg.handlerGetUser))
	routerv1.Post("/feeds", apiCfg.authMiddleware(apiCfg.handlerCreateFeed))
	routerv1.Get("/feeds", apiCfg.HandlerGetFeeds)
	routerv1.Get("/posts", apiCfg.authMiddleware(apiCfg.handlerGetPostsforUser))
	routerv1.Post("/feed_follow", apiCfg.authMiddleware(apiCfg.handlerCreateFeedFollows))
	routerv1.Get("/feed_follow", apiCfg.authMiddleware(apiCfg.handlerGetFeedFollows))
	routerv1.Delete("/feed_follow/{feedfollowID}", apiCfg.authMiddleware(apiCfg.handlerDeleteFeedFollows))

	router.Mount("/v1", routerv1)

	server := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	log.Printf("Server Starting on Port: %v", portString)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PORT: %s\n", portString)
}
