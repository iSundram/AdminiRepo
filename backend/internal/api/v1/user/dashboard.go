
package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

// GetDashboard returns user dashboard data
func (dc *DashboardController) GetDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Dashboard data retrieved successfully",
		"data": gin.H{
			"account_info": gin.H{
				"username": "user123",
				"domain":   "example.com",
				"package":  "Professional",
				"status":   "active",
			},
			"usage": gin.H{
				"disk_used":      "2.5GB",
				"disk_total":     "20GB",
				"bandwidth_used": "45GB",
				"bandwidth_total": "200GB",
				"email_accounts": 15,
				"databases":      3,
				"ftp_accounts":   2,
			},
			"recent_activities": []gin.H{
				{
					"type":      "file_upload",
					"message":   "Uploaded file: image.jpg",
					"timestamp": "2024-01-16T14:30:00Z",
				},
				{
					"type":      "email_created",
					"message":   "Created email account: info@example.com",
					"timestamp": "2024-01-16T10:15:00Z",
				},
			},
			"quick_stats": gin.H{
				"visitors_today":     1250,
				"page_views_today":   3480,
				"emails_sent_today":  45,
				"uptime_percentage":  99.9,
			},
		},
	})
}

// GetQuickActions returns available quick actions
func (dc *DashboardController) GetQuickActions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Quick actions retrieved successfully",
		"data": []gin.H{
			{
				"id":          "file_manager",
				"title":       "File Manager",
				"description": "Manage your website files",
				"icon":        "folder",
				"url":         "/files",
			},
			{
				"id":          "email_accounts",
				"title":       "Email Accounts",
				"description": "Manage email accounts",
				"icon":        "mail",
				"url":         "/email",
			},
			{
				"id":          "databases",
				"title":       "Databases",
				"description": "Manage MySQL databases",
				"icon":        "database",
				"url":         "/databases",
			},
			{
				"id":          "ssl_certificates",
				"title":       "SSL Certificates",
				"description": "Manage SSL certificates",
				"icon":        "shield",
				"url":         "/ssl",
			},
		},
	})
}
