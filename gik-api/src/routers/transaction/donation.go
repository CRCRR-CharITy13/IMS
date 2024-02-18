package transaction

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ListDonationResponse struct {
	Id          int     `json:"ID" binding:"required"`
	CreatedTime string  `json:"createdTime" binding:"required"`
	DonorBy     string  `json:"donorBy" binding:"required"`
	SignedBy    string  `json:"signedBy" binding:"required"`
	TotalValue  float32 `json:"totalValue" binding:"required"`
}

func ListDonations(c *gin.Context) {
	donations := []types.Donation{}

	page := c.Query("page")

	limit := 10
	offset := 0

	if page == "" {
		page = "1"
	}

	date := strings.Split(c.Query("date"), ",")
	user := c.Query("user")

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

	baseQuery := database.Database.Model(&types.Donation{})
	baseQuery = baseQuery.Order("created_at desc")

	if len(date) == 2 && date[0] != "" && date[1] != "" {
		dateStartInt, err := strconv.Atoi(date[0])
		dateStartInt -= 1
		dateEndInt, err := strconv.Atoi(date[1])
		dateEndInt += 86400 // To make sure the filter is inclusive of the entire end date

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid fields - Time",
			})
			return
		}

		dateStart := time.Unix(int64(dateStartInt), 0)
		dateEnd := time.Unix(int64(dateEndInt), 0)

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid fields - Time",
			})
			return
		}

		if err == nil {
			baseQuery = baseQuery.Where("created_at > ?", dateStart)
			baseQuery = baseQuery.Where("created_at < ?", dateEnd)
		}
	}

	userInt, err := strconv.Atoi(user)

	if userInt != 0 {
		// println("user: " + user)
		baseQuery = baseQuery.Where("donor_id = ?", userInt)
	}

	totalCount := int64(0)
	baseQuery.Count(&totalCount)

	baseQuery.Limit(limit).Offset(offset).Find(&donations)

	totalPages := math.Ceil(float64(totalCount) / float64(limit))

	nbDonation := len(donations)
	donationList := make([]ListDonationResponse, nbDonation)
	for idx, donation := range donations {
		database.Database.First(&donation.DonorBy, donation.DonorID)
		database.Database.First(&donation.SignedBy, donation.UserID)
		createdTime := donation.CreatedAt
		strCreatedTime := fmt.Sprintf("%v : %d-%d-%d : %d:%d:%d", createdTime.Weekday(), createdTime.Year(), createdTime.Month(), createdTime.Day(), createdTime.Hour(), createdTime.Minute(), createdTime.Second())
		donationList[idx] = ListDonationResponse{
			Id:          int(donation.ID),
			CreatedTime: strCreatedTime,
			DonorBy:     donation.DonorBy.Name,
			SignedBy:    donation.SignedBy.Username,
			TotalValue:  donation.TotalValue,
		}

	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"data":       donationList,
			"totalPages": totalPages,
		},
	})
}

type AddDonationRequest struct {
	DonorID    int            `json:"donorId" binding:"required"`
	Items      []DonationItem `json:"items" binding:"required"`
	TotalValue float32        `json:"value" binding:"required"`
}

type DonationItem struct {
	SKUName  string `json:"SKUName" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

func AddDonation(c *gin.Context) {
	json := AddDonationRequest{}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
		})
		return
	}

	fmt.Println("Start to create donation")
	fmt.Println(json.Items)

	donation := types.Donation{
		DonorID: uint(json.DonorID),
		UserID:  c.MustGet("userId").(uint),
	}

	// Look-up Donor balance
	donor := types.Donor{}
	database.Database.First(&donor, uint(json.DonorID))
	isSuccess := true
	msgResponse := "Donation created"
	lstDonationItem := []types.DonationItem{}

	for _, inputDonationItem := range json.Items {
		// SKUName: SKU : Name
		// The length of the SKU is 8 character, thus, extract it as follows:
		donationItemSKU := inputDonationItem.SKUName[0:9]
		item := types.Item{}
		baseQuery := database.Database.Model(&types.Item{}).Where("sku = ?", donationItemSKU)
		baseQuery.First(&item)

		donationItem := types.DonationItem{
			DonationID: 0, //dummy value, will be replace after creating the donation
			ItemID:     uint(item.ID),
			Count:      uint(inputDonationItem.Quantity),
		}
		lstDonationItem = append(lstDonationItem, donationItem)
		//database.Database.Create(&donationItem)
		item.StockTotal += inputDonationItem.Quantity
		//
		database.Database.Save(item)

	}

	donation.TotalValue = json.TotalValue

	//TODO: check if this line requried: database.Database.Save(&donation)
	if isSuccess {
		err := database.Database.Create(&donation).Error
		if err != nil {
			c.JSON(500, gin.H{
				"success": false,
				"message": "Could not create donation",
			})
			return
		}
		// save donation items to the database
		for i := 0; i < len(lstDonationItem); i++ {
			//fmt.Println(donation.ID)
			lstDonationItem[i].DonationID = donation.ID
			//fmt.Println(lstDonationItem[i].DonationID)
			database.Database.Create(&lstDonationItem[i])

		}
	}

	c.JSON(200, gin.H{
		"success": isSuccess,
		"message": msgResponse,
	})

}

func DeleteDonation(c *gin.Context) {
	// // get id
	// id := c.Query("id")

	// if id == "" {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Invalid ID",
	// 	})
	// 	return
	// }

	// idInt, err := strconv.Atoi(id)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Invalid ID",
	// 	})
	// 	return
	// }

	// // get donation
	// donation := types.Donation{}
	// database.Database.Where("id = ?", idInt).First(&donation)

	// if donation.ID == 0 {
	// 	c.JSON(400, gin.H{
	// 		"success": false,
	// 		"message": "Donation not found",
	// 	})
	// 	return
	// }

	// // delete all donation items
	// donationItems := []types.DonationItem{}
	// database.Database.Where("donation_id = ?", donation.ID).Delete(&donationItems)

	// // delete transaction
	// database.Database.Delete(&donation)

	// c.JSON(200, gin.H{
	// 	"success": true,
	// 	"message": "Donation deleted",
	// })
}

type donationItemTotalInfo struct {
	ID         uint    `json:"ID" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	SKU        string  `json:"sku" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Quantity   int     `json:"quantity" binding:"required"`
	TotalValue float32 `json:"totalValue" binding:"required" `
}

// Takes an ID query, returns list of donation items (item ID + quantity)
func GetDonationItems(c *gin.Context) {
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

	// Get the donation
	donation := types.Donation{}
	err = database.Database.Where("id = ?", idInt).First(&donation).Error

	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Donation not found",
		})
		return
	}

	// Get the donation items
	donationItems := []types.DonationItem{}
	database.Database.Where("donation_id = ?", donation.ID).Find(&donationItems)

	donationItemsPost := []donationItemTotalInfo{}
	for _, donationItem := range donationItems {

		fmt.Printf("donationItem.ItemID = %d", donationItem.ItemID)
		database.Database.Where("id = ?", donationItem.ItemID).Find(&donationItem.Item)
		fmt.Println("found ----")
		fmt.Printf("name = %s", donationItem.Item.Name)
		donationItemsPost = append(donationItemsPost, donationItemTotalInfo{
			donationItem.ItemID,
			donationItem.Item.Name,
			donationItem.Item.SKU,
			donationItem.Item.Size,
			donationItem.Item.Price,
			int(donationItem.Count),
			donationItem.Item.Price * float32(donationItem.Count),
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    donationItemsPost,
	})
}
