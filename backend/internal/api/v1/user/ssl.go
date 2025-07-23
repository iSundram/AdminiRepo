
package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SSLController struct{}

func NewSSLController() *SSLController {
	return &SSLController{}
}

// GetSSLCertificates returns SSL certificates
func (sc *SSLController) GetSSLCertificates(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "SSL certificates retrieved successfully",
		"data": []gin.H{
			{
				"id":       1,
				"domain":   "example.com",
				"issuer":   "Let's Encrypt",
				"status":   "active",
				"expires":  "2024-06-15T00:00:00Z",
				"auto_renew": true,
				"created":  "2024-01-01T00:00:00Z",
			},
			{
				"id":       2,
				"domain":   "*.example.com",
				"issuer":   "Let's Encrypt",
				"status":   "active",
				"expires":  "2024-07-20T00:00:00Z",
				"auto_renew": true,
				"created":  "2024-01-15T00:00:00Z",
			},
		},
	})
}

// RequestSSLCertificate requests a new SSL certificate
func (sc *SSLController) RequestSSLCertificate(c *gin.Context) {
	var req struct {
		Domain    string `json:"domain" binding:"required"`
		Type      string `json:"type" binding:"required"`
		AutoRenew bool   `json:"auto_renew"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "SSL certificate request submitted successfully",
		"data": gin.H{
			"id":         3,
			"domain":     req.Domain,
			"type":       req.Type,
			"status":     "pending",
			"auto_renew": req.AutoRenew,
		},
	})
}

// RenewSSLCertificate renews an SSL certificate
func (sc *SSLController) RenewSSLCertificate(c *gin.Context) {
	certID := c.Param("id")
	id, err := strconv.Atoi(certID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid certificate ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "SSL certificate renewal initiated",
		"data": gin.H{
			"id":     id,
			"status": "renewing",
		},
	})
}

// DeleteSSLCertificate deletes an SSL certificate
func (sc *SSLController) DeleteSSLCertificate(c *gin.Context) {
	certID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "SSL certificate deleted successfully",
		"data":    gin.H{"id": certID},
	})
}

// InstallCustomSSL installs a custom SSL certificate
func (sc *SSLController) InstallCustomSSL(c *gin.Context) {
	var req struct {
		Domain      string `json:"domain" binding:"required"`
		Certificate string `json:"certificate" binding:"required"`
		PrivateKey  string `json:"private_key" binding:"required"`
		ChainCert   string `json:"chain_cert"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Custom SSL certificate installed successfully",
		"data": gin.H{
			"id":     4,
			"domain": req.Domain,
			"type":   "custom",
			"status": "active",
		},
	})
}
