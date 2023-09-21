package logs

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAdvancedLogs(c *gin.Context) {
	logs := []type_news.AdvancedLog{}

	actionFilter := c.Query("action")
	dateFilter := c.Query("date")
	userFilter := c.Query("user")
	userFilterInt := 0

	dates := []string{
		"",
		"",
	}

	// split dateFilter by comma
	if dateFilter != "" {
		dates = strings.Split(dateFilter, ",")
	}

	if userFilter != "" {
		_, err := strconv.Atoi(userFilter)
		if err != nil {
			// not an int, try to find user id by name
			user := type_news.User{}
			err = database.Database.Where("username = ?", userFilter).Find(&user).Error
			if err != nil {
				c.JSON(400, gin.H{
					"success": false,
					"message": "User doesn't exist",
				})
				return
			}

			userFilterInt = int(user.ID)
		}
	}

	page := c.Query("page")

	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid page",
		})
		return
	}

	// pagination
	limit := 10
	offset := (pageInt - 1) * limit

	baseQuery := database.Database.Model(&type_news.AdvancedLog{})

	if actionFilter != "" {
		baseQuery = baseQuery.Where("path like ?", fmt.Sprintf("%%%s%%", actionFilter))
	}

	if dateFilter != "" {
		baseQuery = baseQuery.Where("timestamp > ?", dates[0]).Where("timestamp < ?", dates[1])
	}

	if userFilter != "" {
		baseQuery = baseQuery.Where(&type_news.AdvancedLog{
			UserID: uint(userFilterInt),
		})
	}

	baseQuery = baseQuery.Order("timestamp desc")

	// get total count
	var totalCount int64
	baseQuery.Count(&totalCount)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	findQuery := baseQuery.Limit(limit).Offset(offset)

	findQuery.Find(&logs)

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":        logs,
			"total":       totalCount,
			"currentPage": pageInt,
			"totalPages":  totalPages,
		},
	})
}
