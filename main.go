package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int32
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "A00", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // 根據主鍵查詢
	println("product" + product.Code)
	db.First(&product, "code = ?", "A00") // 查找Code
	println("product" + product.Code)

	// Update
	db.Model(&product).Update("Price", 200) // 將product更新成200
	// 更新更多
	db.Model(&product).Updates(Product{Price: 200, Code: "Z00"})
	db.Model(&product).Updates(map[string]interface{}{"Price": 250, "Code": "Z99"})

	//Delete
	db.Delete(&product, 1)
}
