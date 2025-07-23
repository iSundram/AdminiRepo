
package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Set Gin mode
    if os.Getenv("GIN_MODE") == "" {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.Default()

    // CORS middleware
    router.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    })

    // Health check endpoint
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status": "ok",
            "service": "AdminiSoftware",
            "version": "1.0.0",
        })
    })

    // API v1 routes
    v1 := router.Group("/api/v1")
    {
        v1.GET("/status", func(c *gin.Context) {
            c.JSON(http.StatusOK, gin.H{
                "message": "AdminiSoftware API v1 is running",
                "website": "https://admini.tech",
                "support": "support@admini.tech",
            })
        })
    }

    // Serve static files for frontend
    router.Static("/static", "./frontend/dist")
    router.StaticFile("/", "./frontend/dist/index.html")

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    log.Printf("AdminiSoftware server starting on 0.0.0.0:%s", port)
    log.Printf("Website: https://admini.tech")
    log.Printf("Support: support@admini.tech")
    
    if err := router.Run("0.0.0.0:" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
