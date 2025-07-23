
package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DatabaseController struct{}

func NewDatabaseController() *DatabaseController {
	return &DatabaseController{}
}

// GetDatabases returns user's databases
func (dc *DatabaseController) GetDatabases(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Databases retrieved successfully",
		"data": []gin.H{
			{
				"id":       1,
				"name":     "user123_main",
				"size":     "15.2MB",
				"tables":   8,
				"charset":  "utf8mb4",
				"collation": "utf8mb4_unicode_ci",
				"created":  "2024-01-01T00:00:00Z",
			},
			{
				"id":       2,
				"name":     "user123_blog",
				"size":     "8.7MB",
				"tables":   12,
				"charset":  "utf8mb4",
				"collation": "utf8mb4_unicode_ci",
				"created":  "2024-01-10T00:00:00Z",
			},
		},
	})
}

// CreateDatabase creates a new database
func (dc *DatabaseController) CreateDatabase(c *gin.Context) {
	var req struct {
		Name      string `json:"name" binding:"required"`
		Charset   string `json:"charset"`
		Collation string `json:"collation"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Database created successfully",
		"data": gin.H{
			"id":        3,
			"name":      req.Name,
			"charset":   req.Charset,
			"collation": req.Collation,
		},
	})
}

// DeleteDatabase deletes a database
func (dc *DatabaseController) DeleteDatabase(c *gin.Context) {
	dbID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Database deleted successfully",
		"data":    gin.H{"id": dbID},
	})
}

// GetDatabaseUsers returns database users
func (dc *DatabaseController) GetDatabaseUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Database users retrieved successfully",
		"data": []gin.H{
			{
				"id":         1,
				"username":   "user123_main",
				"host":       "localhost",
				"privileges": []string{"SELECT", "INSERT", "UPDATE", "DELETE"},
				"created":    "2024-01-01T00:00:00Z",
			},
		},
	})
}

// CreateDatabaseUser creates a new database user
func (dc *DatabaseController) CreateDatabaseUser(c *gin.Context) {
	var req struct {
		Username   string   `json:"username" binding:"required"`
		Password   string   `json:"password" binding:"required,min=8"`
		Host       string   `json:"host"`
		Privileges []string `json:"privileges"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Database user created successfully",
		"data": gin.H{
			"id":         2,
			"username":   req.Username,
			"host":       req.Host,
			"privileges": req.Privileges,
		},
	})
}

// AssignUserToDatabase assigns a user to a database
func (dc *DatabaseController) AssignUserToDatabase(c *gin.Context) {
	var req struct {
		DatabaseID int      `json:"database_id" binding:"required"`
		UserID     int      `json:"user_id" binding:"required"`
		Privileges []string `json:"privileges" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User assigned to database successfully",
		"data": gin.H{
			"database_id": req.DatabaseID,
			"user_id":     req.UserID,
			"privileges":  req.Privileges,
		},
	})
}
