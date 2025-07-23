
package shared

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Login handles user authentication
func (ac *AuthController) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate authentication
	if req.Username == "admin" && req.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Login successful",
			"data": gin.H{
				"token":      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
				"expires_at": time.Now().Add(24 * time.Hour).Unix(),
				"user": gin.H{
					"id":       1,
					"username": "admin",
					"email":    "admin@admini.tech",
					"role":     "admin",
				},
			},
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Invalid credentials",
		})
	}
}

// Logout handles user logout
func (ac *AuthController) Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Logout successful",
	})
}

// RefreshToken refreshes the authentication token
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Token refreshed successfully",
		"data": gin.H{
			"token":      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
			"expires_at": time.Now().Add(24 * time.Hour).Unix(),
		},
	})
}

// GetProfile returns current user profile
func (ac *AuthController) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Profile retrieved successfully",
		"data": gin.H{
			"id":       1,
			"username": "admin",
			"email":    "admin@admini.tech",
			"role":     "admin",
			"created":  "2024-01-01T00:00:00Z",
			"last_login": "2024-01-16T15:30:00Z",
		},
	})
}

// UpdateProfile updates user profile
func (ac *AuthController) UpdateProfile(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Profile updated successfully",
		"data": gin.H{
			"email": req.Email,
		},
	})
}

// Enable2FA enables two-factor authentication
func (ac *AuthController) Enable2FA(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "2FA enabled successfully",
		"data": gin.H{
			"qr_code":    "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...",
			"secret_key": "JBSWY3DPEHPK3PXP",
		},
	})
}

// Disable2FA disables two-factor authentication
func (ac *AuthController) Disable2FA(c *gin.Context) {
	var req struct {
		Code string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "2FA disabled successfully",
	})
}

// VerifyOTP verifies OTP code
func (ac *AuthController) VerifyOTP(c *gin.Context) {
	var req struct {
		Code string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "OTP verified successfully",
	})
}
