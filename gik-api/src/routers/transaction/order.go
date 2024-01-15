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

type AddOrderResponse struct {
	ID           int    `json:"ID" binding:"required"`
	ItemSKUName  string `json:"itemSKUName" binding:"required"`
	LocationName string `json:"locationName" binding:"required"`
	Quantity     int    `json:"quantity" binding:"required"`
}

func AddOrder(c *gin.Context) {
	json := AddOrderRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	fmt.Println("Start to create order")
	fmt.Println(json.Items)

	order := type_news.Order{
		ClientID: uint(json.ClientID),
		UserID:   c.MustGet("userId").(uint),
	}

	// Look-up Client balance
	client := type_news.Client{}
	database.Database.First(&client, uint(json.ClientID))
	totalCost := float32(0)
	isSuccess := true
	msgResponse := "Order created"
	lstAddOrderResponse := []AddOrderResponse{}
	lstOrderItem := []type_news.OrderItem{}

	for _, inputOrderItem := range json.Items {
		// SKUName: SKU : Name
		// The length of the SKU is 8 character, thus, extract it as follows:
		orderItemSKU := inputOrderItem.SKUName[0:9]
		item := type_news.Item{}
		baseQuery := database.Database.Model(&type_news.Item{}).Where("sku = ?", orderItemSKU)
		baseQuery.First(&item)

		if item.StockTotal < inputOrderItem.Quantity {
			isSuccess = false
			msgResponse = fmt.Sprintf("There are only %d of %s\n", item.StockTotal, inputOrderItem.SKUName)
			break
		}

		totalCost += float32(inputOrderItem.Quantity) * item.Price
		if totalCost > client.Balance {
			isSuccess = false
			msgResponse = fmt.Sprintf("Current balance (%f) is not enough", client.Balance)
			break
		}
		// TODO: update the current item.StockTotal
		//item.StockTotal -= product.Quantity

		//database.Database.Save(item)

		// create order item

		orderItem := type_news.OrderItem{
			OrderID: 0, //dummy value, will be replace after creating the order
			ItemID:  uint(item.ID),
			Count:   inputOrderItem.Quantity,
		}
		lstOrderItem = append(lstOrderItem, orderItem)
		//database.Database.Create(&orderItem)
	}

	order.TotalCost = totalCost

	//TODO: check if this line requried: database.Database.Save(&order)
	if isSuccess {
		err := database.Database.Create(&order).Error
		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": "Could not create order",
			})
			return
		}
		// save order items to the database
		for i := 0; i < len(lstOrderItem); i++ {
			//fmt.Println(order.ID)
			lstOrderItem[i].OrderID = order.ID
			//fmt.Println(lstOrderItem[i].OrderID)
			database.Database.Create(&lstOrderItem[i])

		}
		// remove items from locatons
		idx := 0
		for _, orderItem := range lstOrderItem {
			database.Database.First(&orderItem.Item, orderItem.ItemID)
			// get list of item's warehouses
			database.Database.Preload("Warehouses").Where("item_id = ?", orderItem.Item.ID).Find(&orderItem.Item.Warehouses)
			// fmt.Printf("for item %s \n", orderItem.Item.Name)
			// fmt.Println(orderItem.Item.Warehouses)
			remainQtt := orderItem.Count
			var removeQtt int
			strItemSKUName := orderItem.Item.SKU + " : " + orderItem.Item.Name
			for _, warehouse := range orderItem.Item.Warehouses {
				if remainQtt > warehouse.Stock {
					removeQtt = warehouse.Stock
				} else {
					removeQtt = remainQtt
				}
				var location type_news.Location
				database.Database.First(&location, warehouse.LocationID)
				idx++
				tmpAddOrderResponse := AddOrderResponse{
					ID:           idx,
					ItemSKUName:  strItemSKUName,
					LocationName: location.Name,
					Quantity:     removeQtt,
				}
				lstAddOrderResponse = append(lstAddOrderResponse, tmpAddOrderResponse)
				// remove these amount of items from the location
				warehouse.Stock -= removeQtt
				// if remain stock > 0 : save; else: delete current record
				if warehouse.Stock > 0 {
					database.Database.Save(warehouse)
				} else {
					database.Database.Delete(&warehouse)
				}
				//
				remainQtt -= removeQtt
				if remainQtt == 0 {
					break
				}
			}
			if remainQtt > 0 {
				idx++
				tmpAddOrderResponse := AddOrderResponse{
					ID:           idx,
					ItemSKUName:  strItemSKUName,
					LocationName: "Common Storage",
					Quantity:     remainQtt,
				}
				lstAddOrderResponse = append(lstAddOrderResponse, tmpAddOrderResponse)
			}
			orderItem.Item.StockTotal -= orderItem.Count
		}

		//update Client.Balance
		client.Balance -= totalCost
		database.Database.Save(client)
	}

	c.JSON(200, gin.H{
		"success": isSuccess,
		"message": msgResponse,
		"data":    lstAddOrderResponse,
	})

}

func DeleteOrder(c *gin.Context) {
	// get id
	id := c.Query("id")

	if id == "" {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
		})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - ID",
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
			"message": "Invalid fields - ID",
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
