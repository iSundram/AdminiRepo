
package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmailController struct{}

func NewEmailController() *EmailController {
	return &EmailController{}
}

// GetEmailAccounts returns email accounts
func (ec *EmailController) GetEmailAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email accounts retrieved successfully",
		"data": []gin.H{
			{
				"id":       1,
				"email":    "admin@example.com",
				"quota":    "1000MB",
				"used":     "250MB",
				"status":   "active",
				"created":  "2024-01-01T00:00:00Z",
			},
			{
				"id":       2,
				"email":    "info@example.com",
				"quota":    "500MB",
				"used":     "125MB",
				"status":   "active",
				"created":  "2024-01-05T00:00:00Z",
			},
		},
	})
}

// CreateEmailAccount creates a new email account
func (ec *EmailController) CreateEmailAccount(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
		Quota    int    `json:"quota" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Email account created successfully",
		"data": gin.H{
			"id":     3,
			"email":  req.Email,
			"quota":  req.Quota,
			"status": "active",
		},
	})
}

// UpdateEmailAccount updates an email account
func (ec *EmailController) UpdateEmailAccount(c *gin.Context) {
	emailID := c.Param("id")
	id, err := strconv.Atoi(emailID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email ID"})
		return
	}

	var req struct {
		Password string `json:"password"`
		Quota    int    `json:"quota"`
		Status   string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email account updated successfully",
		"data": gin.H{
			"id":     id,
			"quota":  req.Quota,
			"status": req.Status,
		},
	})
}

// DeleteEmailAccount deletes an email account
func (ec *EmailController) DeleteEmailAccount(c *gin.Context) {
	emailID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email account deleted successfully",
		"data":    gin.H{"id": emailID},
	})
}

// GetEmailForwarders returns email forwarders
func (ec *EmailController) GetEmailForwarders(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email forwarders retrieved successfully",
		"data": []gin.H{
			{
				"id":          1,
				"source":      "contact@example.com",
				"destination": "admin@example.com",
				"status":      "active",
				"created":     "2024-01-10T00:00:00Z",
			},
		},
	})
}

// CreateEmailForwarder creates a new email forwarder
func (ec *EmailController) CreateEmailForwarder(c *gin.Context) {
	var req struct {
		Source      string `json:"source" binding:"required,email"`
		Destination string `json:"destination" binding:"required,email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Email forwarder created successfully",
		"data": gin.H{
			"id":          2,
			"source":      req.Source,
			"destination": req.Destination,
			"status":      "active",
		},
	})
}
