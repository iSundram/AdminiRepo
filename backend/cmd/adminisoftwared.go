
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// Admin controllers
	adminUsers "adminisoftware/internal/api/v1/admin"
	
	// User controllers  
	userControllers "adminisoftware/internal/api/v1/user"
	
	// Reseller controllers
	resellerControllers "adminisoftware/internal/api/v1/reseller"
	
	// Shared controllers
	sharedControllers "adminisoftware/internal/api/v1/shared"
	
	// AI controllers
	aiControllers "adminisoftware/internal/api/v1/ai"
	
	// Plugin controllers
	pluginControllers "adminisoftware/internal/api/v1/plugins"
	
	// Middleware
	"adminisoftware/internal/api/v1/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Gin router
	r := gin.Default()

	// Add middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RateLimitMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "healthy",
			"version": "1.0.0",
			"service": "AdminiSoftware Control Panel",
		})
	})

	// Initialize controllers
	authController := sharedControllers.NewAuthController()
	
	// Admin controllers
	adminUserController := adminUsers.NewUserController()
	adminAccountController := adminUsers.NewAccountController()
	adminPackageController := adminUsers.NewPackageController()
	adminSystemController := adminUsers.NewSystemController()
	
	// User controllers
	userDashboardController := userControllers.NewDashboardController()
	userDomainController := userControllers.NewDomainController()
	userFileController := userControllers.NewFileController()
	userEmailController := userControllers.NewEmailController()
	userDatabaseController := userControllers.NewDatabaseController()
	userSSLController := userControllers.NewSSLController()
	
	// Reseller controllers
	resellerDashboardController := resellerControllers.NewDashboardController()
	resellerAccountController := resellerControllers.NewAccountController()
	
	// AI controllers
	aiAssistantController := aiControllers.NewAssistantController()
	
	// Plugin controllers
	pluginMarketplaceController := pluginControllers.NewMarketplaceController()

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Authentication routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/logout", authController.Logout)
			auth.POST("/refresh", authController.RefreshToken)
			auth.GET("/profile", middleware.AuthMiddleware(), authController.GetProfile)
			auth.PUT("/profile", middleware.AuthMiddleware(), authController.UpdateProfile)
			auth.POST("/2fa/enable", middleware.AuthMiddleware(), authController.Enable2FA)
			auth.POST("/2fa/disable", middleware.AuthMiddleware(), authController.Disable2FA)
			auth.POST("/2fa/verify", middleware.AuthMiddleware(), authController.VerifyOTP)
		}

		// Admin routes (protected)
		admin := v1.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			// User management
			adminUsers := admin.Group("/users")
			{
				adminUsers.GET("", adminUserController.GetUsers)
				adminUsers.POST("", adminUserController.CreateUser)
				adminUsers.PUT("/:id", adminUserController.UpdateUser)
				adminUsers.DELETE("/:id", adminUserController.DeleteUser)
			}

			// Account management
			adminAccounts := admin.Group("/accounts")
			{
				adminAccounts.GET("", adminAccountController.GetAccounts)
				adminAccounts.POST("", adminAccountController.CreateAccount)
				adminAccounts.PUT("/:id/suspend", adminAccountController.SuspendAccount)
				adminAccounts.PUT("/:id/unsuspend", adminAccountController.UnsuspendAccount)
				adminAccounts.DELETE("/:id", adminAccountController.DeleteAccount)
			}

			// Package management
			adminPackages := admin.Group("/packages")
			{
				adminPackages.GET("", adminPackageController.GetPackages)
				adminPackages.POST("", adminPackageController.CreatePackage)
				adminPackages.PUT("/:id", adminPackageController.UpdatePackage)
				adminPackages.DELETE("/:id", adminPackageController.DeletePackage)
			}

			// System management
			adminSystem := admin.Group("/system")
			{
				adminSystem.GET("/info", adminSystemController.GetSystemInfo)
				adminSystem.GET("/logs", adminSystemController.GetSystemLogs)
				adminSystem.POST("/services/:service/restart", adminSystemController.RestartService)
				adminSystem.POST("/update", adminSystemController.UpdateSystem)
				adminSystem.GET("/backups", adminSystemController.GetBackups)
				adminSystem.POST("/backups", adminSystemController.CreateBackup)
			}
		}

		// Reseller routes (protected)
		reseller := v1.Group("/reseller")
		reseller.Use(middleware.AuthMiddleware(), middleware.ResellerMiddleware())
		{
			// Dashboard
			reseller.GET("/dashboard", resellerDashboardController.GetDashboard)
			reseller.GET("/statistics", resellerDashboardController.GetAccountStatistics)

			// Client account management
			resellerAccounts := reseller.Group("/accounts")
			{
				resellerAccounts.GET("", resellerAccountController.GetAccounts)
				resellerAccounts.POST("", resellerAccountController.CreateAccount)
				resellerAccounts.PUT("/:id", resellerAccountController.UpdateAccount)
				resellerAccounts.PUT("/:id/suspend", resellerAccountController.SuspendAccount)
				resellerAccounts.PUT("/:id/unsuspend", resellerAccountController.UnsuspendAccount)
				resellerAccounts.DELETE("/:id", resellerAccountController.DeleteAccount)
			}
		}

		// User panel routes (protected)
		user := v1.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			// Dashboard
			user.GET("/dashboard", userDashboardController.GetDashboard)
			user.GET("/quick-actions", userDashboardController.GetQuickActions)

			// Domain management
			userDomains := user.Group("/domains")
			{
				userDomains.GET("", userDomainController.GetDomains)
				userDomains.POST("", userDomainController.AddDomain)
				userDomains.DELETE("/:id", userDomainController.DeleteDomain)
				userDomains.GET("/:id/dns", userDomainController.GetDNSRecords)
				userDomains.POST("/:id/dns", userDomainController.AddDNSRecord)
			}

			// File management
			userFiles := user.Group("/files")
			{
				userFiles.GET("", userFileController.GetFiles)
				userFiles.POST("/upload", userFileController.UploadFile)
				userFiles.POST("/folder", userFileController.CreateFolder)
				userFiles.DELETE("", userFileController.DeleteFile)
				userFiles.PUT("/rename", userFileController.RenameFile)
				userFiles.GET("/content", userFileController.GetFileContent)
				userFiles.POST("/content", userFileController.SaveFileContent)
			}

			// Email management
			userEmail := user.Group("/email")
			{
				userEmail.GET("/accounts", userEmailController.GetEmailAccounts)
				userEmail.POST("/accounts", userEmailController.CreateEmailAccount)
				userEmail.PUT("/accounts/:id", userEmailController.UpdateEmailAccount)
				userEmail.DELETE("/accounts/:id", userEmailController.DeleteEmailAccount)
				userEmail.GET("/forwarders", userEmailController.GetEmailForwarders)
				userEmail.POST("/forwarders", userEmailController.CreateEmailForwarder)
			}

			// Database management
			userDatabases := user.Group("/databases")
			{
				userDatabases.GET("", userDatabaseController.GetDatabases)
				userDatabases.POST("", userDatabaseController.CreateDatabase)
				userDatabases.DELETE("/:id", userDatabaseController.DeleteDatabase)
				userDatabases.GET("/users", userDatabaseController.GetDatabaseUsers)
				userDatabases.POST("/users", userDatabaseController.CreateDatabaseUser)
				userDatabases.POST("/assign", userDatabaseController.AssignUserToDatabase)
			}

			// SSL management
			userSSL := user.Group("/ssl")
			{
				userSSL.GET("", userSSLController.GetSSLCertificates)
				userSSL.POST("", userSSLController.RequestSSLCertificate)
				userSSL.POST("/:id/renew", userSSLController.RenewSSLCertificate)
				userSSL.DELETE("/:id", userSSLController.DeleteSSLCertificate)
				userSSL.POST("/custom", userSSLController.InstallCustomSSL)
			}
		}

		// AI Assistant routes (protected)
		ai := v1.Group("/ai")
		ai.Use(middleware.AuthMiddleware())
		{
			ai.POST("/assistant", aiAssistantController.GetAssistantResponse)
			ai.GET("/analysis", aiAssistantController.GetSystemAnalysis)
			ai.GET("/suggestions", aiAssistantController.GetOptimizationSuggestions)
			ai.GET("/predictions", aiAssistantController.PredictResourceUsage)
		}

		// Plugin marketplace routes (protected)
		plugins := v1.Group("/plugins")
		plugins.Use(middleware.AuthMiddleware())
		{
			plugins.GET("", pluginMarketplaceController.GetPlugins)
			plugins.GET("/:id", pluginMarketplaceController.GetPlugin)
			plugins.POST("/:id/install", pluginMarketplaceController.InstallPlugin)
			plugins.DELETE("/:id", pluginMarketplaceController.UninstallPlugin)
			plugins.GET("/installed", pluginMarketplaceController.GetInstalledPlugins)
			plugins.POST("/:id/update", pluginMarketplaceController.UpdatePlugin)
		}
	}

	// Static file serving for frontend
	r.Static("/assets", "./frontend/dist/assets")
	r.StaticFile("/", "./frontend/dist/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("AdminiSoftware server starting on 0.0.0.0:%s", port)
	log.Println("Website: https://admini.tech")
	log.Println("Support: support@admini.tech")
	
	if err := r.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
