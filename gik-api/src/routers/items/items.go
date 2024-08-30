package items

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	//"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
)

// Function to create a new item
// CreateItem is the function to create a new item

// START CREATE ITEM PROD VERSION
func GetCreditValueIfExist(c *gin.Context) {
	var createNewItemRequest struct {
		CategoryCode     string `json:"category_code" binding:"required"`
		SubCategory1Code int    `json:"subcategory1_code" binding:"required"`
		SubCategory2Code string `json:"subcategory2_code" binding:"required"`
		DescriptionCode  int    `json:"description_code" binding:"required"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&createNewItemRequest); err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid fields - Missing values"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	db := database.Database

	// Verify the codes exist in the database
	if err := CheckCode(db, createNewItemRequest.CategoryCode, createNewItemRequest.SubCategory1Code, createNewItemRequest.SubCategory2Code, createNewItemRequest.DescriptionCode); err != nil {
		c.JSON(400, gin.H{"success": false, "message": "One or more codes are invalid"})
		log.Printf("Error verifying codes: %v", err)
		return
	}

	// Check if item family already exists
	var itemFamilyID int
	var creditValue float64
	query := "SELECT item_family_id, current_credit_value FROM items_families WHERE category_code = ? AND subcategory1_code = ? AND subcategory2_code = ? AND description_code = ?"
	err := db.QueryRow(query, createNewItemRequest.CategoryCode, createNewItemRequest.SubCategory1Code, createNewItemRequest.SubCategory2Code, createNewItemRequest.DescriptionCode).Scan(&itemFamilyID, &creditValue)

	if err == nil {
		// Item family exists, return the credit value and item family ID
		c.JSON(200, gin.H{"success": true, "message": "Item family exists", "data": gin.H{
			"item_family_id": itemFamilyID,
			"credit_value":   creditValue,
		}})
		log.Printf("Item family exists: %d, credit value: %f", itemFamilyID, creditValue)
		return
	} else if err != sql.ErrNoRows {
		// Some other error occurred
		c.JSON(500, gin.H{"success": false, "message": "Database error checking item family"})
		log.Printf("Error checking item family: %v", err)
		return
	}

	// No item family exists, return a response to allow new item family creation
	c.JSON(200, gin.H{"success": true, "message": "No item family exists", "data": gin.H{
		"item_family_id": nil,
	}})
}

func CreateItemAndItemFamily(c *gin.Context) {
	var createNewItemRequest struct {
		CategoryCode     string  `json:"category_code" binding:"required"`
		SubCategory1Code int     `json:"subcategory1_code" binding:"required"`
		SubCategory2Code string  `json:"subcategory2_code" binding:"required"`
		DescriptionCode  int     `json:"description_code" binding:"required"`
		Category         string  `json:"category_name,omitempty"`
		SubCategory1     string  `json:"subcategory1_name,omitempty"`
		SubCategory2     string  `json:"subcategory2_name,omitempty"`
		Description      string  `json:"description_name,omitempty"`
		Size             string  `json:"item_size" binding:"required"`
		CreditValue      float64 `json:"current_credit_value"`
		Quantity         int     `json:"item_total_quantity" binding:"required"`
		ItemFamilyID     int     `json:"item_family_id"`
	}

	// Bind JSON input to struct
	if err := c.ShouldBindJSON(&createNewItemRequest); err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid fields - Missing values"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	db := database.Database

	// Step 1: Verify that categoryCode, subCategory1Code, subCategory2Code, and descriptionCode exist
	if err := CheckCode(db, createNewItemRequest.CategoryCode, createNewItemRequest.SubCategory1Code, createNewItemRequest.SubCategory2Code, createNewItemRequest.DescriptionCode); err != nil {
		c.JSON(400, gin.H{"success": false, "message": "One or more codes are invalid"})
		log.Printf("Code existence check failed: %v", err)
		return
	}

	// Step 2: Fetch names if they are not provided
	if createNewItemRequest.Category == "" || createNewItemRequest.SubCategory1 == "" || createNewItemRequest.SubCategory2 == "" || createNewItemRequest.Description == "" {
		category, subCategory1Name, subCategory2Name, descriptionName, err := GetNamesByCodes(
			db,
			createNewItemRequest.CategoryCode,
			createNewItemRequest.SubCategory1Code,
			createNewItemRequest.SubCategory2Code,
			createNewItemRequest.DescriptionCode,
		)
		if err != nil {
			c.JSON(500, gin.H{"success": false, "message": "Failed to retrieve names based on codes"})
			log.Printf("Error retrieving names: %v", err)
			return
		}

		if createNewItemRequest.Category == "" {
			createNewItemRequest.Category = category
		}
		if createNewItemRequest.SubCategory1 == "" {
			createNewItemRequest.SubCategory1 = subCategory1Name
		}
		if createNewItemRequest.SubCategory2 == "" {
			createNewItemRequest.SubCategory2 = subCategory2Name
		}
		if createNewItemRequest.Description == "" {
			createNewItemRequest.Description = descriptionName
		}
	}

	// Step 3: Check if the ItemFamily exists or create a new one
	if createNewItemRequest.ItemFamilyID == 0 {
		var itemFamilyID int
		query := "SELECT item_family_id FROM items_families WHERE category_code = ? AND subcategory1_code = ? AND subcategory2_code = ? AND description_code = ?"
		err := db.QueryRow(query, createNewItemRequest.CategoryCode, createNewItemRequest.SubCategory1Code, createNewItemRequest.SubCategory2Code, createNewItemRequest.DescriptionCode).Scan(&itemFamilyID)

		if err == sql.ErrNoRows {
			// No existing ItemFamily, create a new one
			itemSKU := createNewItemRequest.CategoryCode + strconv.Itoa(createNewItemRequest.SubCategory1Code) + createNewItemRequest.SubCategory2Code + strconv.Itoa(createNewItemRequest.DescriptionCode)
			itemName := createNewItemRequest.Category + " " + createNewItemRequest.SubCategory1 + " " + createNewItemRequest.SubCategory2 + " " + createNewItemRequest.Description
			now := time.Now().Format(time.RFC3339)

			result, err := db.Exec("INSERT INTO items_families (category_code, subcategory1_code, subcategory2_code, description_code, current_credit_value, item_name, item_sku, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
				createNewItemRequest.CategoryCode, createNewItemRequest.SubCategory1Code, createNewItemRequest.SubCategory2Code, createNewItemRequest.DescriptionCode, createNewItemRequest.CreditValue, itemName, itemSKU, now, now)
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to create item family"})
				log.Printf("Error creating item family: %v", err)
				return
			}

			lastInsertID, err := result.LastInsertId()
			if err != nil {
				c.JSON(500, gin.H{"success": false, "message": "Failed to retrieve new item family ID"})
				log.Printf("Error retrieving new item family ID: %v", err)
				return
			}
			createNewItemRequest.ItemFamilyID = int(lastInsertID)
		} else if err != nil {
			// Some other error occurred
			c.JSON(500, gin.H{"success": false, "message": "Error checking item family"})
			log.Printf("Error checking item family: %v", err)
			return
		} else {
			// ItemFamily exists, use the existing ID
			createNewItemRequest.ItemFamilyID = itemFamilyID
		}
	}

	// Step 4: Check if the item with the given size already exists for this ItemFamily
	var existingItemID int
	query := "SELECT item_family_id FROM items WHERE item_family_id = ? AND item_size = ?"
	err := db.QueryRow(query, createNewItemRequest.ItemFamilyID, createNewItemRequest.Size).Scan(&existingItemID)
	if err == nil {
		// Item already exists, return an error
		c.JSON(400, gin.H{"success": false, "message": "Item size already exists for this item family"})
		return
	} else if err != sql.ErrNoRows {
		// Some other error occurred
		c.JSON(500, gin.H{"success": false, "message": "Error checking existing item"})
		log.Printf("Error checking existing item: %v", err)
		return
	}

	// Step 5: Create the new item associated with the item family
	now := time.Now().Format(time.RFC3339)
	_, err = db.Exec("INSERT INTO items (item_family_id, item_size, item_total_quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		createNewItemRequest.ItemFamilyID, createNewItemRequest.Size, createNewItemRequest.Quantity, now, now)
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": "Failed to create item"})
		log.Printf("Error creating item: %v", err)
		return
	}

	c.JSON(200, gin.H{"success": true, "message": "Item created successfully", "item_family_id": createNewItemRequest.ItemFamilyID})
}

// END CREATE ITEM PROD VERSION

// START CREATE ITEM DEV VERSION
type CreateItemInput struct {
	Category     string  `json:"category_name" binding:"required"`
	SubCategory1 string  `json:"subcategory1_name" binding:"required"`
	SubCategory2 string  `json:"subcategory2_name" binding:"required"`
	Description  string  `json:"description_name" binding:"required"`
	Size         string  `json:"item_size" binding:"required"`
	CreditValue  float64 `json:"current_credit_value"`
	Quantity     int     `json:"item_total_quantity"`
}

func CreateItem(c *gin.Context) {
	var createNewItemRequest CreateItemInput

	if err := c.ShouldBindJSON(&createNewItemRequest); err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid fields - Missing values"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	if createNewItemRequest.CreditValue == 0 {
		c.JSON(400, gin.H{"success": false, "message": "Credit value cannot be null"})
		return
	}

	db := database.Database

	var categoryCode, subCategory2Code, sizeCode string
	var subCategory1Code, descriptionCode int

	// Check if category exists
	log.Println("Checking category")
	err := db.QueryRow("SELECT category_code FROM categories WHERE category_name = ?", createNewItemRequest.Category).Scan(&categoryCode)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid category name"})
		log.Printf("Error checking category: %v", err)
		return
	}
	log.Printf("Category found: %s", categoryCode)

	// Check if subcategory 1 exists
	log.Println("Checking subCategory1")
	err = db.QueryRow("SELECT subcategory1_code FROM subcategories1 WHERE subcategory1_name = ?", createNewItemRequest.SubCategory1).Scan(&subCategory1Code)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid subCategory1 name"})
		log.Printf("Error checking subCategory1: %v", err)
		return
	}
	log.Printf("SubCategory1 found: %d", subCategory1Code)

	// Check if subcategory 2 exists
	log.Println("Checking subCategory2")
	err = db.QueryRow("SELECT subcategory2_code FROM subCategories2 WHERE subcategory2_name = ?", createNewItemRequest.SubCategory2).Scan(&subCategory2Code)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid subCategory2 name"})
		log.Printf("Error checking subCategory2: %v", err)
		return
	}
	log.Printf("SubCategory2 found: %s", subCategory2Code)

	// Check if description exists
	log.Println("Checking description")
	err = db.QueryRow("SELECT description_code FROM items_descriptions WHERE description_name = ?", createNewItemRequest.Description).Scan(&descriptionCode)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid description name"})
		log.Printf("Error checking description: %v", err)
		return
	}
	log.Printf("Description found: %d", descriptionCode)

	// Check if size exists
	log.Println("Checking size")
	err = db.QueryRow("SELECT item_size FROM sizes WHERE item_size = ?", createNewItemRequest.Size).Scan(&sizeCode)
	if err != nil {
		c.JSON(400, gin.H{"success": false, "message": "Invalid size name"})
		log.Printf("Error checking size: %v", err)
		return
	}
	log.Printf("Size found: %s", sizeCode)

	// Check if item family already exists
	var existingItemFamily types.ItemFamily
	query := "SELECT item_family_id FROM items_families WHERE category_code = ? AND subcategory1_code = ? AND subcategory2_code = ? AND description_code = ?"
	log.Println("Checking item family")
	errItemFamilyExist := db.QueryRow(query, categoryCode, subCategory1Code, subCategory2Code, descriptionCode).Scan(&existingItemFamily.ItemFamilyID)

	if errItemFamilyExist == nil {
		log.Printf("Item family exists: %d", existingItemFamily.ItemFamilyID)

		/* Get the credit value and send it to front.
		* That need another module.
		* Another route to get the credit value.

		 */

		// Item family exists, check if the size exists in items
		var existingItemSize types.Item
		query = "SELECT item_family_id FROM items WHERE item_family_id = ? AND item_size = ?"
		log.Println("Checking item size in existing item family")
		err = db.QueryRow(query, existingItemFamily.ItemFamilyID, sizeCode).Scan(&existingItemSize.ItemFamilyID)
		if err == nil {
			c.JSON(400, gin.H{"success": false, "message": "Item Family and Item size already exists in item table"})
			log.Printf("Item size already exists for item table: %d", existingItemFamily.ItemFamilyID)
			return
		}
	} else if errItemFamilyExist != sql.ErrNoRows {
		c.JSON(500, gin.H{"success": false, "message": "Database error checking item family"})
		log.Printf("Error checking item family: %v", errItemFamilyExist)
		return
	}

	if createNewItemRequest.CreditValue <= 0 {
		c.JSON(400, gin.H{"success": false, "message": "Credit value must be greater than 0"})
		log.Println("Invalid credit value")
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": "Failed to start transaction"})
		log.Printf("Error starting transaction: %v", err)
		return
	}

	now := time.Now().Format(time.RFC3339)
	var itemSKU, itemName string
	itemSKU = categoryCode + strconv.Itoa(subCategory1Code) + subCategory2Code + strconv.Itoa(descriptionCode)
	itemName = createNewItemRequest.SubCategory1 + " " + createNewItemRequest.SubCategory2 + " " + createNewItemRequest.Description

	queryInsertItemTable := "INSERT INTO items (item_family_id, item_size, item_total_quantity, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"

	// Create new item when item family already exists and item size doesn't exist yet
	if errItemFamilyExist == nil {
		log.Println("Inserting new item")
		_, err = tx.Exec(queryInsertItemTable, existingItemFamily.ItemFamilyID, sizeCode, createNewItemRequest.Quantity, now, now)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"success": false, "message": "Failed to insert item where Item family already exists"})
			log.Printf("Error inserting item: %v", err)
			return
		}
		tx.Commit()
		c.JSON(200, gin.H{"success": true, "message": "Item created successfully"})
		return
	}

	// Create new item family
	log.Println("Inserting new item family")
	result, err := tx.Exec("INSERT INTO items_families (category_code, subcategory1_code, subcategory2_code, description_code, current_credit_value, item_sku, item_name, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", categoryCode, subCategory1Code, subCategory2Code, descriptionCode, createNewItemRequest.CreditValue, itemSKU, itemName, now, now)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"success": false, "message": "Failed to insert item family"})
		log.Printf("Error inserting item family: %v", err)
		return
	}

	// Get new item family ID
	itemFamilyID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"success": false, "message": "Failed to get item family ID"})
		log.Printf("Error getting item family ID: %v", err)
		return
	}
	log.Printf("New item family ID: %d", itemFamilyID)

	// Create new item
	log.Println("Inserting new item in new item family")
	_, err = tx.Exec(queryInsertItemTable, itemFamilyID, sizeCode, createNewItemRequest.Quantity, now, now)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"success": false, "message": "Failed to insert item"})
		log.Printf("Error inserting item: %v", err)
		return
	}

	// Create new item credit value history
	log.Println("Inserting item credit value history")
	_, err = tx.Exec("INSERT INTO items_credit_value_history (item_family_id, date, credit_value) VALUES (?, ?, ?)", itemFamilyID, createNewItemRequest.CreditValue, now)
	if err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"success": false, "message": "Failed to insert item credit value history"})
		log.Printf("Error inserting item credit value history: %v", err)
		return
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		c.JSON(500, gin.H{"success": false, "message": "Failed to commit transaction"})
		log.Printf("Error committing transaction: %v", err)
		return
	}

	c.JSON(200, gin.H{"success": true, "message": "Item, Item Family and Item Credit Value History created successfully"})
	log.Println("Item created successfully")
}

// END CREATE ITEM DEV VERSION

type UpdateItemInput struct {
	ItemFamilyID    int    `json:"item_family_id" binding:"required"`
	Size            string `json:"item_size" binding:"required"`
	CreditValue     int    `json:"current_credit_value"`
	Quantity        int    `json:"item_total_quantity"`
	JustifyQuantity string `json:"justify_quantity"`
}

func Edititem(c *gin.Context) {
	var UpdateItemInput UpdateItemInput

	if err := c.ShouldBindJSON(&UpdateItemInput); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Missing values",
			"error":   err.Error(),
		})
		return
	}

	// get itemQuantity by id and size
	var itemTotalQuantity int

	err := database.Database.QueryRow("SELECT item_total_quantity FROM items WHERE item_family_id = ? AND item_size = ?", UpdateItemInput.ItemFamilyID, UpdateItemInput.Size).Scan(&itemTotalQuantity)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Item  doesn't exist",
		})
		return
	}
	fmt.Printf(" Item total quantity: %d", itemTotalQuantity)

	// get credit value from item family by id
	var creditValue int
	err = database.Database.QueryRow("SELECT current_credit_value FROM items_families WHERE item_family_id = ?", UpdateItemInput.ItemFamilyID).Scan(&creditValue)
	if err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields - Item family doesn't exist",
		})
		return
	}
	fmt.Printf("Credit Value: %d", creditValue)

	// check if creditValue ou quantity have not changed

	if (UpdateItemInput.CreditValue == creditValue) && (UpdateItemInput.CreditValue > 0) && (UpdateItemInput.Quantity == itemTotalQuantity) {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Nothing to update",
		})
		return
	}

	// star transaction
	tx, err := database.Database.Begin()
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Failed to start transaction",
		})
		return
	}

	now := time.Now().Format(time.RFC3339)

	fmt.Printf("Check if quantity changed: %d", UpdateItemInput.Quantity)
	fmt.Printf("Quantity changed: %t", UpdateItemInput.Quantity != itemTotalQuantity)
	// check if quantity changed
	if UpdateItemInput.Quantity != itemTotalQuantity {
		fmt.Print("Quantity changed")
		fmt.Printf("Check if justify quantity exists: %s", UpdateItemInput.JustifyQuantity)
		// check if justify quantity exists
		if UpdateItemInput.JustifyQuantity == "" {
			c.JSON(400, gin.H{
				"success": false,
				"message": "Invalid fields - Missing justify quantity",
			})
			return
		}
		fmt.Printf("Justify quantity exists: %s", UpdateItemInput.JustifyQuantity)

		fmt.Printf("Update item total quantity: %d", UpdateItemInput.Quantity)
		// update item total quantity
		_, err = tx.Exec("UPDATE items SET item_total_quantity = ?, updated_at = ? WHERE item_family_id = ? AND item_size = ?", UpdateItemInput.Quantity, now, UpdateItemInput.ItemFamilyID, UpdateItemInput.Size)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{
				"success": false,
				"message": "Failed to update item total quantity",
			})
			return
		}
		fmt.Printf("Item total quantity updated: %d", UpdateItemInput.Quantity)

		fmt.Printf("Create Item Quantity Log: %d ...", UpdateItemInput.Quantity)
		// add log for quantity changed
		timestamps := time.Now().Unix()
		clientIP := c.ClientIP()
		fmt.Printf("Client IP: %s", clientIP)

		action := "Quantity changed form " + strconv.Itoa(itemTotalQuantity) + " to " + strconv.Itoa(UpdateItemInput.Quantity) + " on item " + strconv.Itoa(UpdateItemInput.ItemFamilyID) + " by user 1 " //+ strconv.Itoa(c.MustGet("userId").(uint))

		_, err = tx.Exec("INSERT INTO simple_logs (action, timestamp, ip_address, user_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", action, timestamps, clientIP, 1, now, now)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{
				"success": false,
				"message": "Failed to add log",
				"error":   err,
			})
			return
		}

		c.JSON(200, gin.H{
			"success": true,
			"message": "Item quantity updated successfully",
		})
	}

	// check if credit value changed and if so, update credit value
	if (UpdateItemInput.CreditValue != creditValue) && (UpdateItemInput.CreditValue > 0) {

		// update credit value in item family
		fmt.Printf("Update credit value: %d", UpdateItemInput.CreditValue)
		_, err = tx.Exec("UPDATE items_families SET current_credit_value = ?, updated_at = ? WHERE item_family_id = ?", UpdateItemInput.CreditValue, now, UpdateItemInput.ItemFamilyID)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{
				"success": false,
				"message": "Failed to update credit value",
				"error":   err,
			})
			return
		}
		fmt.Printf("Credit value updated: %d", UpdateItemInput.CreditValue)

		fmt.Printf("Create Item Credit Value History: %d ...", UpdateItemInput.CreditValue)
		// create credit value in item credit value history
		_, err = tx.Exec("INSERT INTO items_credit_value_history (item_family_id, date, credit_value) VALUES (?, ?, ?)", UpdateItemInput.ItemFamilyID, now, UpdateItemInput.CreditValue)
		if err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{
				"success": false,
				"message": "Failed to insert credit value in item credit value history",
			})
			return
		}
		fmt.Printf("Item credit value history created: %d", UpdateItemInput.CreditValue)

		// send notification to all users with credit value changed
		fmt.Println("In construction - Notification epic part : Send notification to all users with credit value changed")

		c.JSON(200, gin.H{
			"success": true,
			"message": "Credit value updated successfully",
		})
	}

	tx.Commit()
	c.JSON(200, gin.H{"success": true, "message": "Updated successfully"})
	fmt.Printf("Item created successfully")
}

type GetItemFamily struct {
	ItemFamilyID       int              `json:"item_family_id"`
	ItemName           sql.NullString   `json:"item_name"`
	ItemSKU            sql.NullString   `json:"item_sku"`
	CurrentCreditValue int              `json:"current_credit_value"`
	Items              []GetItemSummary `json:"items"`
}

type GetItemSummary struct {
	ItemSize          string `json:"item_size"`
	ItemTotalQuantity int    `json:"item_total_quantity"`
}

func GetItems(c *gin.Context) {
	row, err := database.Database.Query(`
    SELECT 
        f.item_family_id,
				f.item_sku,
				f.item_name,
				f.current_credit_value,
        i.item_size,
				i.item_total_quantity
    FROM 
        items_families f
    INNER JOIN 
        items i ON f.item_family_id = i.item_family_id
`)

	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	// Map pour regrouper les familles avec leurs items
	familiesMap := make(map[int]*GetItemFamily)
	for row.Next() {
		var family GetItemFamily
		var item GetItemSummary

		err = row.Scan(&family.ItemFamilyID, &family.ItemSKU, &family.ItemName, &family.CurrentCreditValue, &item.ItemSize, &item.ItemTotalQuantity)

		if err != nil {
			log.Fatal(err)
		}
		// Si la famille existe déjà dans la map, ajouter l'item à la liste
		if existingFamily, exists := familiesMap[family.ItemFamilyID]; exists {
			existingFamily.Items = append(existingFamily.Items, item)
		} else {
			// Sinon, créer une nouvelle famille et ajouter l'item
			family.Items = []GetItemSummary{item}
			familiesMap[family.ItemFamilyID] = &family
		}
	}

	// Convertir la map en slice pour la réponse JSON
	var families []GetItemFamily
	for _, family := range familiesMap {
		families = append(families, *family)
	}

	totalFamilies := len(families)
	c.JSON(200, gin.H{"success": true, "data": gin.H{
		"itemsFamilies": families,
		"total":         totalFamilies,
	}})
}

// HELPERS PART

// CheckCode checks if category, subcategory1, subcategory2 and description codes are valid
func CheckCode(db *sql.DB, categoryCode string, subCategory1Code int, subCategory2Code string, descriptionCode int) error {
	// Check if category exists
	var temp string
	err := db.QueryRow("SELECT category_code FROM categories WHERE category_code = ?", categoryCode).Scan(&temp)
	if err != nil {
		log.Printf("Error checking category code: %v", err)
		return fmt.Errorf("invalid category code")
	}

	// Check if subcategory 1 exists
	err = db.QueryRow("SELECT subcategory1_code FROM subcategories1 WHERE subcategory1_code = ?", subCategory1Code).Scan(&temp)
	if err != nil {
		log.Printf("Error checking subCategory1 code: %v", err)
		return fmt.Errorf("invalid subCategory1 code")
	}

	// Check if subcategory 2 exists
	err = db.QueryRow("SELECT subcategory2_code FROM subcategories2 WHERE subcategory2_code = ?", subCategory2Code).Scan(&temp)
	if err != nil {
		log.Printf("Error checking subCategory2 code: %v", err)
		return fmt.Errorf("invalid subCategory2 code")
	}

	// Check if description exists
	err = db.QueryRow("SELECT description_code FROM items_descriptions WHERE description_code = ?", descriptionCode).Scan(&temp)
	if err != nil {
		log.Printf("Error checking description code: %v", err)
		return fmt.Errorf("invalid description code")
	}

	return nil
}

func GetNamesByCodes(db *sql.DB, categoryCode string, subCategory1Code int, subCategory2Code string, descriptionCode int) (string, string, string, string, error) {
	var category, subCategory1Name, subCategory2Name, descriptionName string

	// Fetch Category name by its code
	err := db.QueryRow("SELECT category_name FROM categories WHERE category_code = ?", categoryCode).Scan(&category)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error fetching Category name: %v", err)
	}

	// Fetch SubCategory1 name by its code
	err = db.QueryRow("SELECT subcategory1_name FROM subcategories1 WHERE subcategory1_code = ?", subCategory1Code).Scan(&subCategory1Name)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error fetching SubCategory1 name: %v", err)
	}

	// Fetch SubCategory2 name by its code
	err = db.QueryRow("SELECT subcategory2_name FROM subcategories2 WHERE subcategory2_code = ?", subCategory2Code).Scan(&subCategory2Name)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error fetching SubCategory2 name: %v", err)
	}

	// Fetch Description name by its code
	err = db.QueryRow("SELECT description_name FROM items_descriptions WHERE description_code = ?", descriptionCode).Scan(&descriptionName)
	if err != nil {
		return "", "", "", "", fmt.Errorf("error fetching Description name: %v", err)
	}

	return category, subCategory1Name, subCategory2Name, descriptionName, nil
}
