
package reseller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

// GetAccounts returns reseller's client accounts
func (ac *AccountController) GetAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Client accounts retrieved successfully",
		"data": []gin.H{
			{
				"id":       1,
				"domain":   "client1.com",
				"username": "client1",
				"package":  "Professional",
				"status":   "active",
				"created":  "2024-01-01T00:00:00Z",
				"revenue":  "$29.99",
			},
			{
				"id":       2,
				"domain":   "client2.com",
				"username": "client2",
				"package":  "Basic",
				"status":   "active",
				"created":  "2024-01-05T00:00:00Z",
				"revenue":  "$19.99",
			},
		},
	})
}

// CreateAccount creates a new client account
func (ac *AccountController) CreateAccount(c *gin.Context) {
	var req struct {
		Domain   string `json:"domain" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Package  string `json:"package" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Client account created successfully",
		"data": gin.H{
			"id":       3,
			"domain":   req.Domain,
			"username": req.Username,
			"email":    req.Email,
			"package":  req.Package,
			"status":   "active",
		},
	})
}

// UpdateAccount updates a client account
func (ac *AccountController) UpdateAccount(c *gin.Context) {
	accountID := c.Param("id")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	var req struct {
		Package string `json:"package"`
		Status  string `json:"status"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Client account updated successfully",
		"data": gin.H{
			"id":      id,
			"package": req.Package,
			"status":  req.Status,
		},
	})
}

// SuspendAccount suspends a client account
func (ac *AccountController) SuspendAccount(c *gin.Context) {
	accountID := c.Param("id")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Client account suspended successfully",
		"data":    gin.H{"id": id, "status": "suspended"},
	})
}

// UnsuspendAccount unsuspends a client account
func (ac *AccountController) UnsuspendAccount(c *gin.Context) {
	accountID := c.Param("id")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Client account unsuspended successfully",
		"data":    gin.H{"id": id, "status": "active"},
	})
}

// DeleteAccount deletes a client account
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Client account deleted successfully",
		"data":    gin.H{"id": accountID},
	})
}
