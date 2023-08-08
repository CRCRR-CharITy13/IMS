package items

import (
	"GIK_Web/database"
	"GIK_Web/types"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"os"
	"strings"
)

type export struct {
	ID       uint    `gorm:"primaryKey;autoIncrement:false" csv:"-"`
	Name     string  `json:"name"`
	SKU      string  `json:"name"`
	Category string  `json:"category"` //Gender and such
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	Size     string  `json:"size"`
}

func ExportItems(c *gin.Context) {
	name := c.Query("name")
	sku := c.Query("sku")
	tags := strings.Split(c.Query("tags"), ",")

	items := []export{}

	baseQuery := database.Database.Model(&types.Item{})

	baseQuery = baseQuery.Order("sku, FIELD(size, 'XXL',  'XL', 'L', 'M', 'S', 'XS', 'XXS'), size")

	for _, tag := range tags {
		baseQuery = baseQuery.Where("category LIKE ?", "%"+tag+"%")
	}
	if name != "" {
		baseQuery = baseQuery.Where("name LIKE ?", "%"+name+"%")
	}
	if sku != "" {
		baseQuery = baseQuery.Where("sku LIKE ?", "%"+sku+"%")
	}

	baseQuery.Find(&items)

	fileName := uuid.New().String() + ".csv"
	csvFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	err = gocsv.MarshalFile(&items, csvFile)
	c.File(fileName)

	os.Remove(fileName)
}
