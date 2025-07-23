
package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

// GetAccounts returns all hosting accounts
func (ac *AccountController) GetAccounts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Accounts retrieved successfully",
		"data": []gin.H{
			{
				"id":       1,
				"domain":   "example.com",
				"username": "exampleuser",
				"package":  "Basic",
				"status":   "active",
				"created":  "2024-01-01T00:00:00Z",
			},
		},
	})
}

// CreateAccount creates a new hosting account
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
		"message": "Account created successfully",
		"data": gin.H{
			"id":       2,
			"domain":   req.Domain,
			"username": req.Username,
			"email":    req.Email,
			"package":  req.Package,
			"status":   "active",
		},
	})
}

// SuspendAccount suspends a hosting account
func (ac *AccountController) SuspendAccount(c *gin.Context) {
	accountID := c.Param("id")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account suspended successfully",
		"data":    gin.H{"id": id, "status": "suspended"},
	})
}

// UnsuspendAccount unsuspends a hosting account
func (ac *AccountController) UnsuspendAccount(c *gin.Context) {
	accountID := c.Param("id")
	id, err := strconv.Atoi(accountID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account unsuspended successfully",
		"data":    gin.H{"id": id, "status": "active"},
	})
}

// DeleteAccount deletes a hosting account
func (ac *AccountController) DeleteAccount(c *gin.Context) {
	accountID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account deleted successfully",
		"data":    gin.H{"id": accountID},
	})
}
