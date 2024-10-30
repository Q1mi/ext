# the extention of golang time.Time

A type alias for time.Time, with some custom method.
- UnmarshalJSON/MarshalJSON
- GormDataType/GormDBDataType
- Value
- Scan

## Usage


```golang
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Q1mi/ext"
)

// gin GORM simple demo

func initDB() (*gorm.DB, error) {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

type Product struct {
	Name     string   `json:"name"`
	Birthday ext.Time `json:"birthday"` // use ext.Time instead of time.Time
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		var p Product
		if err := c.ShouldBindJSON(&p); err != nil {
			log.Printf("ShouldBindJSON fail, err:%v\n", err)
			c.JSON(400, gin.H{"err": err})
			return
		}
		db.Create(p)
		var pv Product
		if err := db.Model(&Product{}).First(&pv).Error; err != nil {
			log.Printf("db.First query fail, err:%v\n", err)
			c.JSON(400, gin.H{"err": err})
			return
		}

		c.JSON(200, gin.H{
			"data": pv,
		})
	})

	r.Run()
}
```
