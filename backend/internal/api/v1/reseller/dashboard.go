
package reseller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DashboardController struct{}

func NewDashboardController() *DashboardController {
	return &DashboardController{}
}

// GetDashboard returns reseller dashboard data
func (dc *DashboardController) GetDashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Reseller dashboard data retrieved successfully",
		"data": gin.H{
			"account_info": gin.H{
				"username":      "reseller123",
				"company":       "Hosting Solutions Inc",
				"status":        "active",
				"accounts_used": 45,
				"accounts_limit": 100,
			},
			"usage": gin.H{
				"disk_used":      "125GB",
				"disk_total":     "500GB",
				"bandwidth_used": "850GB",
				"bandwidth_total": "2TB",
				"accounts_active": 42,
				"accounts_suspended": 3,
			},
			"revenue": gin.H{
				"monthly_revenue":    "$2,850.00",
				"total_revenue":      "$34,200.00",
				"pending_payments":   "$450.00",
				"commission_rate":    "15%",
			},
			"recent_activities": []gin.H{
				{
					"type":      "account_created",
					"message":   "Created account: newclient.com",
					"timestamp": "2024-01-16T14:30:00Z",
				},
				{
					"type":      "payment_received",
					"message":   "Payment received: $29.99",
					"timestamp": "2024-01-16T10:15:00Z",
				},
			},
		},
	})
}

// GetAccountStatistics returns account statistics
func (dc *DashboardController) GetAccountStatistics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Account statistics retrieved successfully",
		"data": gin.H{
			"total_accounts":     45,
			"active_accounts":    42,
			"suspended_accounts": 3,
			"accounts_by_package": gin.H{
				"Basic":        20,
				"Professional": 18,
				"Premium":      7,
			},
			"monthly_growth": gin.H{
				"new_accounts":    8,
				"cancelled_accounts": 2,
				"growth_rate":     "13.3%",
			},
		},
	})
}
