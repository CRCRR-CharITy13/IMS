package transaction

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ListTransactions(c *gin.Context) {
	transactions := []types.Transaction{}

	page := c.Query("page")

	limit := 10
	offset := 0

	if page == "" {
		page = "1"
	}

	type_ := c.Query("type")
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

	baseQuery := database.Database.Model(&types.Transaction{})
	baseQuery = baseQuery.Order("timestamp desc")
	// println("type: " + type_)
	if type_ == "export" {
		// println("type is export")
		baseQuery = baseQuery.Where("type = ?", 0)
	} else if type_ == "import" {
		// println("type is import")
		baseQuery = baseQuery.Where("type = ?", 1)
	} else {
		// println("type is all")
	}
	if len(date) == 2 && date[0] != "" && date[1] != "" {
		dateStartInt, err := strconv.Atoi(date[0])
		dateStartInt -= 1
		dateEndInt, err := strconv.Atoi(date[1])
		dateEndInt += 86400 // To make sure the filter is inclusive of the entire end date
		if err == nil {
			// print("date start: ")
			// println(dateStartInt)
			// print("date end: ")
			// println(dateEndInt)
			baseQuery = baseQuery.Where("timestamp > ?", dateStartInt)
			baseQuery = baseQuery.Where("timestamp < ?", dateEndInt)
		}
	}

	userInt, err := strconv.Atoi(user)

	if userInt != 0 {
		// println("user: " + user)
		baseQuery = baseQuery.Where("client_id = ?", userInt)
	}

	totalCount := int64(0)
	baseQuery.Count(&totalCount)

	baseQuery.Limit(limit).Offset(offset).Find(&transactions)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":       transactions,
			"totalPages": totalPages,
		},
	})
}

type addRequest struct {
	ClientID int           `json:"clientId" binding:"required"`
	Type     bool          `json:"type"`
	Products []productBody `json:"products" binding:"required"`
}

type productBody struct {
	ID       int `json:"id" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

func AddTransaction(c *gin.Context) {
	json := addRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid request",
		})
		return
	}

	totalQuantity := 0

	transaction := types.Transaction{
		Type:      json.Type,
		ClientID:  uint(json.ClientID),
		SignerID:  c.MustGet("userId").(uint),
		Timestamp: time.Now().Unix(),
	}

	database.Database.Create(&transaction)

	for _, product := range json.Products {
		// get items
		item := types.Item{}
		baseQuery := database.Database.Model(&types.Item{}).Where("id = ?", product.ID)
		baseQuery.First(&item)

		if json.Type {
			item.Quantity += product.Quantity
		} else {
			item.Quantity -= product.Quantity
		}

		database.Database.Save(item)

		// create transaction item
		transactionItem := types.TransactionItem{
			TransactionID: transaction.ID,
			Quantity:      product.Quantity,
			ProductID:     product.ID,
		}

		database.Database.Create(&transactionItem)

		totalQuantity += product.Quantity
	}

	transaction.TotalQuantity = totalQuantity

	database.Database.Save(&transaction)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Transaction created",
	})
}

func DeleteTransaction(c *gin.Context) {
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

	// get transaction
	transaction := types.Transaction{}
	database.Database.Where("id = ?", idInt).First(&transaction)

	if transaction.ID == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Transaction not found",
		})
		return
	}

	// delete all transaction items
	transactionItems := []types.TransactionItem{}
	database.Database.Where("transaction_id = ?", transaction.ID).Delete(&transactionItems)

	// delete transaction
	database.Database.Delete(&transaction)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Transaction deleted",
	})
}

type itemsPost struct {
	ID         int     `json:"ID"`
	Name       string  `json:"name"`
	SKU        string  `json:"sku"`
	Size       string  `json:"size"`
	Price      float32 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalValue float32 `json:"totalValue"`
}

func GetTransactionItems(c *gin.Context) {
	// id
	id := c.Query("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid ID",
		})
		return
	}

	// get transaction
	transaction := types.Transaction{}
	database.Database.Where("id = ?", idInt).First(&transaction)

	if transaction.ID == 0 {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Transaction not found",
		})
		return
	}

	// get transaction items
	transactionItems := []types.TransactionItem{}
	database.Database.Where("transaction_id = ?", transaction.ID).Find(&transactionItems)

	transactionItemsInfo := []types.Item{}
	for _, item := range transactionItems {
		itemInfo := types.Item{}
		database.Database.Where("id = ?", item.ID).Find(&itemInfo)
		transactionItemsInfo = append(transactionItemsInfo, itemInfo)
	}

	itemCount := len(transactionItems)

	transactionItemsPost := []itemsPost{}

	for i := 0; i < itemCount; i++ {
		transactionItemsPost = append(transactionItemsPost, itemsPost{
			transactionItems[i].ProductID,
			transactionItemsInfo[i].Name,
			transactionItemsInfo[i].SKU,
			transactionItemsInfo[i].Size,
			transactionItemsInfo[i].Price,
			transactionItems[i].Quantity,
			transactionItemsInfo[i].Price * float32(transactionItems[i].Quantity),
		})
		//transactionItemsPost[i].ID = transactionItems[i].ProductID
		//transactionItemsPost[i].Name = transactionItemsInfo[i].Name
		//transactionItemsPost[i].SKU = transactionItemsInfo[i].SKU
		//transactionItemsPost[i].Price = transactionItemsInfo[i].Price
		//transactionItemsPost[i].Quantity = transactionItems[i].Quantity
		//transactionItemsPost[i].TotalValue = float32(transactionItemsPost[i].Quantity) * transactionItemsPost[i].Price
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    transactionItemsPost,
	})
}
