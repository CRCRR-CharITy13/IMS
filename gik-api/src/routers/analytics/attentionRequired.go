package analytics

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID            string
	SKU           string  `json:"sku"`
	ProductID     uint    `json:"id"`
	Price         float32 `json:"price"`
	StockQuantity float32 `json:"stockQuantity"`
	Name          string  `json:"name"`
	Category      string  `json:"category"`
	Gender        string  `json:"gender"`
	Season        string  `json:"season"`
	Location      string  `json:"location"`
	Image         []byte  `json:"image"`
}

type post struct {
	ID       uint    `json:"id"`
	SKU      string  `json:"sku"`
	Name     string  `json:"name"`
	Size     string  `json:"size"`
	Quantity float32 `json:"quantity"`
}

const (
	max = 3
)

func AttentionRequired(c *gin.Context) {

	var urgentItems []post
	// database.Database.Raw("SELECT * FROM `AP6_wc_product_meta_lookup` WHERE `stock_quantity` < ?", max).Find(&postData)

	limit := 5
	offset := 0

	page := c.Query("page")

	if page == "" {
		page = "1"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid page number",
		})
		return
	}

	offset = (pageInt - 1) * limit

	baseQuery := database.Database.Model(&types.Item{}).Where("quantity < ?", max).Where("name is not null").Order("quantity desc").Order("sku desc").Order("size desc")

	itemCount := int64(0)
	baseQuery.Count(&itemCount)

	totalPages := int(math.Ceil(float64(itemCount) / float64(limit)))

	err = baseQuery.Limit(limit).Offset(offset).Find(&urgentItems).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Unable to get data",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":       urgentItems,
			"totalPages": totalPages,
		},
	})

}
