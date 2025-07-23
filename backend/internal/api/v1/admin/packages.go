
package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PackageController struct{}

func NewPackageController() *PackageController {
	return &PackageController{}
}

// GetPackages returns all hosting packages
func (pc *PackageController) GetPackages(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Packages retrieved successfully",
		"data": []gin.H{
			{
				"id":          1,
				"name":        "Basic",
				"disk_quota":  5000,
				"bandwidth":   50000,
				"email_quota": 100,
				"ftp_quota":   5,
				"db_quota":    5,
				"price":       9.99,
				"features": []string{
					"SSL Certificate",
					"Email Accounts",
					"File Manager",
				},
			},
			{
				"id":          2,
				"name":        "Professional",
				"disk_quota":  20000,
				"bandwidth":   200000,
				"email_quota": 500,
				"ftp_quota":   25,
				"db_quota":    25,
				"price":       29.99,
				"features": []string{
					"SSL Certificate",
					"Email Accounts",
					"File Manager",
					"Database Access",
					"Backup Service",
				},
			},
		},
	})
}

// CreatePackage creates a new hosting package
func (pc *PackageController) CreatePackage(c *gin.Context) {
	var req struct {
		Name        string   `json:"name" binding:"required"`
		DiskQuota   int      `json:"disk_quota" binding:"required"`
		Bandwidth   int      `json:"bandwidth" binding:"required"`
		EmailQuota  int      `json:"email_quota" binding:"required"`
		FTPQuota    int      `json:"ftp_quota" binding:"required"`
		DBQuota     int      `json:"db_quota" binding:"required"`
		Price       float64  `json:"price" binding:"required"`
		Features    []string `json:"features"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Package created successfully",
		"data": gin.H{
			"id":          3,
			"name":        req.Name,
			"disk_quota":  req.DiskQuota,
			"bandwidth":   req.Bandwidth,
			"email_quota": req.EmailQuota,
			"ftp_quota":   req.FTPQuota,
			"db_quota":    req.DBQuota,
			"price":       req.Price,
			"features":    req.Features,
		},
	})
}

// UpdatePackage updates an existing package
func (pc *PackageController) UpdatePackage(c *gin.Context) {
	packageID := c.Param("id")
	id, err := strconv.Atoi(packageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid package ID"})
		return
	}

	var req struct {
		Name        string   `json:"name"`
		DiskQuota   int      `json:"disk_quota"`
		Bandwidth   int      `json:"bandwidth"`
		EmailQuota  int      `json:"email_quota"`
		FTPQuota    int      `json:"ftp_quota"`
		DBQuota     int      `json:"db_quota"`
		Price       float64  `json:"price"`
		Features    []string `json:"features"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Package updated successfully",
		"data": gin.H{
			"id":          id,
			"name":        req.Name,
			"disk_quota":  req.DiskQuota,
			"bandwidth":   req.Bandwidth,
			"email_quota": req.EmailQuota,
			"ftp_quota":   req.FTPQuota,
			"db_quota":    req.DBQuota,
			"price":       req.Price,
			"features":    req.Features,
		},
	})
}

// DeletePackage deletes a package
func (pc *PackageController) DeletePackage(c *gin.Context) {
	packageID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Package deleted successfully",
		"data":    gin.H{"id": packageID},
	})
}
