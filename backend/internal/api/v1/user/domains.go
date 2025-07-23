
package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DomainController struct{}

func NewDomainController() *DomainController {
	return &DomainController{}
}

// GetDomains returns user's domains
func (dc *DomainController) GetDomains(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Domains retrieved successfully",
		"data": []gin.H{
			{
				"id":          1,
				"domain":      "example.com",
				"type":        "main",
				"status":      "active",
				"ssl_status":  "active",
				"created":     "2024-01-01T00:00:00Z",
				"expires":     "2025-01-01T00:00:00Z",
			},
			{
				"id":          2,
				"domain":      "subdomain.example.com",
				"type":        "subdomain",
				"status":      "active",
				"ssl_status":  "none",
				"created":     "2024-01-15T00:00:00Z",
			},
		},
	})
}

// AddDomain adds a new domain/subdomain
func (dc *DomainController) AddDomain(c *gin.Context) {
	var req struct {
		Domain string `json:"domain" binding:"required"`
		Type   string `json:"type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Domain added successfully",
		"data": gin.H{
			"id":     3,
			"domain": req.Domain,
			"type":   req.Type,
			"status": "active",
		},
	})
}

// DeleteDomain removes a domain
func (dc *DomainController) DeleteDomain(c *gin.Context) {
	domainID := c.Param("id")
	id, err := strconv.Atoi(domainID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid domain ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Domain deleted successfully",
		"data":    gin.H{"id": id},
	})
}

// GetDNSRecords returns DNS records for a domain
func (dc *DomainController) GetDNSRecords(c *gin.Context) {
	domainID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "DNS records retrieved successfully",
		"data": gin.H{
			"domain_id": domainID,
			"records": []gin.H{
				{
					"id":    1,
					"name":  "@",
					"type":  "A",
					"value": "192.168.1.100",
					"ttl":   3600,
				},
				{
					"id":    2,
					"name":  "www",
					"type":  "CNAME",
					"value": "example.com",
					"ttl":   3600,
				},
				{
					"id":    3,
					"name":  "@",
					"type":  "MX",
					"value": "mail.example.com",
					"priority": 10,
					"ttl":   3600,
				},
			},
		},
	})
}

// AddDNSRecord adds a new DNS record
func (dc *DomainController) AddDNSRecord(c *gin.Context) {
	domainID := c.Param("id")

	var req struct {
		Name     string `json:"name" binding:"required"`
		Type     string `json:"type" binding:"required"`
		Value    string `json:"value" binding:"required"`
		TTL      int    `json:"ttl"`
		Priority int    `json:"priority"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "DNS record added successfully",
		"data": gin.H{
			"domain_id": domainID,
			"record": gin.H{
				"id":       4,
				"name":     req.Name,
				"type":     req.Type,
				"value":    req.Value,
				"ttl":      req.TTL,
				"priority": req.Priority,
			},
		},
	})
}
