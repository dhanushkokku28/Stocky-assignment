package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "stock-reward-system/internal/db"
    "stock-reward-system/internal/api"
)

func main() {
    // Initialize the database
    err := db.Connect()
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    }

    // Set up Gin router
    router := gin.Default()

    // Set up routes
    api.RegisterRoutes(router)

    // Start the server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}