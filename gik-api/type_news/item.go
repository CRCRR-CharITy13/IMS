type Items struct {
	gorm.Model
	SKU        string
	Name       string
	StockTotal int
	Size       string
	Price      float32
}