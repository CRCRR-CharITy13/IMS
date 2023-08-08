/*
invoice creates invoices for transactions.

Takes a transaction and the related items and generates an invoice with all the
relevant information.
*/

package invoice

import (
	"GIK_Web/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	generator "github.com/angelodlfrtr/go-invoice-generator"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type items struct {
	ID         int     `json:"ID"`
	ItemName   string  `json:"name" binding:"required"`
	Size       string  `json:"size" binding:"required"`
	SKU        string  `json:"sku" binding:"required"`
	Price      float32 `json:"price" binding:"required"`
	Quantity   int     `json:"quantity" binding:"required"`
	TotalValue float32 `json:"totalValue"`
}

type data struct {
	Data []items `json:"data" binding:"required"`

	ClientName string `json:"name" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

func GetInvoice(c *gin.Context) {
	// Get the request
	json := data{}

	//
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Invalid fields",
		})
		return
	}

	c.JSON(400, gin.H{
		"data": json,
	})

	doc, _ := generator.New(generator.Invoice, &generator.Options{
		CurrencySymbol:  " ",
		TextTypeInvoice: "Invoice",
		AutoPrint:       true,
	})

	doc.SetHeader(&generator.HeaderFooter{
		Text:       "<right>Gifts in Kind.</right>",
		Pagination: true,
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>THIS INVOICE IS FOR INFORMATION. YOU DON’T NEED TO MAKE ANY PAYMENT.</center>",
		Pagination: true,
	})

	doc.SetRef("GIK")
	doc.SetVersion("1.0")

	// Make adjustable?
	doc.SetNotes("<b>Customer Notes<b><br>Let me know when ready to pick up. Thank you.")
	doc.SetDate(time.Now().Format("2006/01/01"))

	logoBytes, _ := ioutil.ReadFile("assets/Logo.png")

	doc.SetCompany(&generator.Contact{
		Name: "Gifts In Kind",
		Logo: &logoBytes,
		Address: &generator.Address{
			Address:    "CRC Rideau-Rockcliffe CRC Unit 3 - 815 St. Laurent Blvd",
			PostalCode: "K1K 3A7",
			City:       "Ottawa, Ontario",
			Country:    " ",
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: json.ClientName,
		Address: &generator.Address{
			Address:    json.Address,
			PostalCode: "",
			City:       "",
			Country:    "",
		},
	})

	for _, i := range json.Data {
		doc.AppendItem(&generator.Item{
			Name:        i.ItemName,
			Description: "   " + i.Size + "\nSKU: " + i.SKU,
			UnitCost:    fmt.Sprintf("%f", i.Price),
			Quantity:    strconv.Itoa(i.Quantity),
		})
	}

	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
	}

	name := uuid.New().String() + ".pdf"

	err = pdf.OutputFileAndClose(name)

	c.File(name)

	os.Remove(name)

	if err != nil {
		log.Fatal(err)
	}

	utils.CreateSimpleLog(c, "Generated an invoice")
}
