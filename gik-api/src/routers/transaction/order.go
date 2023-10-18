package transaction

import (
	"GIK_Web/database"
	"GIK_Web/type_news"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ListOrderResponse struct {
	Id          int     `json:"ID" binding:"required"`
	CreatedTime string  `json:"createdTime" binding:"required"`
	ClientName  string  `json:"clientName" binding:"required"`
	SignedBy    string  `json:"signedBy" binding:"required"`
	TotalCost   float32 `json:"totalCost" binding:"required"`
}

func ListOrders(c *gin.Context) {
	orders := []type_news.Order{}

	page := c.Query("page")

	limit := 10
	offset := 0

	if page == "" {
		page = "1"
	}

	date := strings.Split(c.Query("date"), ",")
	user := c.Query("user")
	// print("date: ")
	// for _, d := range date {
	// 	println(" " + d + " ")
	// }
	// println()
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid page",
		})
		return
	}

	// pagination
	offset = (pageInt - 1) * limit

	baseQuery := database.Database.Model(&type_news.Order{})
	baseQuery = baseQuery.Order("created_at desc")

	if len(date) == 2 && date[0] != "" && date[1] != "" {
		dateStartInt, err := strconv.Atoi(date[0])
		dateStartInt -= 1
		dateEndInt, err := strconv.Atoi(date[1])
		dateEndInt += 86400 // To make sure the filter is inclusive of the entire end date
		if err == nil {
			baseQuery = baseQuery.Where("created_at > ?", dateStartInt)
			baseQuery = baseQuery.Where("created_at < ?", dateEndInt)
		}
	}

	userInt, err := strconv.Atoi(user)

	if userInt != 0 {
		// println("user: " + user)
		baseQuery = baseQuery.Where("client_id = ?", userInt)
	}

	totalCount := int64(0)
	baseQuery.Count(&totalCount)

	baseQuery.Limit(limit).Offset(offset).Find(&orders)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	nbOrder := len(orders)
	orderList := make([]ListOrderResponse, nbOrder)
	for idx, order := range orders {
		database.Database.First(&order.Client, order.ClientID)
		database.Database.First(&order.SignedBy, order.UserID)
		createdTime := order.CreatedAt
		strCreatedTime := fmt.Sprintf("%v : %d-%d-%d : %d:%d:%d", createdTime.Weekday(), createdTime.Year(), createdTime.Month(), createdTime.Day(), createdTime.Hour(), createdTime.Minute(), createdTime.Second())
		orderList[idx] = ListOrderResponse{
			Id:          int(order.ID),
			CreatedTime: strCreatedTime,
			ClientName:  order.Client.OrgName,
			SignedBy:    order.SignedBy.Username,
			TotalCost:   order.TotalCost,
		}

	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":       orderList,
			"totalPages": totalPages,
		},
	})
}

type AddOrderRequest struct {
	ClientID int         `json:"clientId" binding:"required"`
	Items    []OrderItem `json:"items" binding:"required"`
}

type OrderItem struct {
	SKUName  string `json:"SKUName" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func AddOrder(c *gin.Context) {
	json := AddOrderRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request",
		})
		return
	}

	totalCost := float32(0)

	order := type_news.Order{
		ClientID: uint(json.ClientID),
		UserID:   c.MustGet("userId").(uint),
	}

	err := database.Database.Create(&order).Error
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Could not create order",
		})
		return
	}
	fmt.Println("Start to create order")
	fmt.Println(json.Items)
	for _, inputOrderItem := range json.Items {
		// SKUName: SKU : Name
		// The length of the SKU is 8 character, thus, extract it as follows:
		orderItemSKU := inputOrderItem.SKUName[0:9]
		item := type_news.Item{}
		baseQuery := database.Database.Model(&type_news.Item{}).Where("sku = ?", orderItemSKU)
		baseQuery.First(&item)

		// TODO: update the current item.StockTotal
		//item.StockTotal -= product.Quantity

		//database.Database.Save(item)

		// create transaction item
		orderItem := type_news.OrderItem{
			OrderID: order.ID,
			ItemID:  uint(item.ID),
			Count:   inputOrderItem.Quantity,
		}

		database.Database.Create(&orderItem)

		totalCost += float32(inputOrderItem.Quantity) * item.Price
	}

	order.TotalCost = totalCost

	database.Database.Save(&order)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Order created",
	})
}

func DeleteOrder(c *gin.Context) {
	// get id
	id := c.Query("id")

	if id == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	// get order
	order := type_news.Order{}
	database.Database.Where("id = ?", idInt).First(&order)

	if order.ID == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Order not found",
		})
		return
	}

	// delete all order items
	orderItems := []type_news.OrderItem{}
	database.Database.Where("order_id = ?", order.ID).Delete(&orderItems)

	// delete transaction
	database.Database.Delete(&order)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Order deleted",
	})
}

type orderItemTotalInfo struct {
	ID        uint    `json:"ID" binding:"required"`
	Name      string  `json:"name" binding:"required"`
	SKU       string  `json:"sku" binding:"required"`
	Size      string  `json:"size" binding:"required"`
	Price     float32 `json:"price" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required"`
	TotalCost float32 `json:"totalCost" binding:"required" `
}

// Takes an ID query, returns list of order items (item ID + quantity)
func GetOrderItems(c *gin.Context) {
	// Get ID from query
	id := c.Query("id")

	// Check that the ID is an int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	// Get the order
	order := type_news.Order{}
	err = database.Database.Where("id = ?", idInt).First(&order).Error

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Order not found",
		})
		return
	}

	// Get the order items
	orderItems := []type_news.OrderItem{}
	database.Database.Where("order_id = ?", order.ID).Find(&orderItems)

	orderItemsPost := []orderItemTotalInfo{}
	for _, orderItem := range orderItems {

		fmt.Printf("orderItem.ItemID = %d", orderItem.ItemID)
		database.Database.Where("id = ?", orderItem.ItemID).Find(&orderItem.Item)
		fmt.Println("found ----")
		fmt.Printf("name = %s", orderItem.Item.Name)
		orderItemsPost = append(orderItemsPost, orderItemTotalInfo{
			orderItem.ItemID,
			orderItem.Item.Name,
			orderItem.Item.SKU,
			orderItem.Item.Size,
			orderItem.Item.Price,
			orderItem.Count,
			orderItem.Item.Price * float32(orderItem.Count),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    orderItemsPost,
	})
}
