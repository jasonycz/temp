package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// User user model
type User struct {
	ID              int64
	Name            string
	Email           string
	EmailVerifiedAt string
	Password        string
	RememberToken   string
	CreatedAt       string
	UpdatedAt       string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/eshop?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//自动检查 User 结构是否变化，变化则进行迁移
	// db.AutoMigrate(&User{})

	// 增
	// db.Create(&User{Code: "L1212", Price: 1000})

	// 查
	var user User
	db.First(&user, 1) // 找到id为1的产品

	fmt.Printf("%+v", user)
	// db.First(&User, "code = ?", "L1212") // 找出 code 为 l1212 的产品

	// 改 - 更新产品的价格为 2000
	// db.Model(&User).Update("Price", 2000)

	// 删 - 删除产品
	// db.Delete(&User)
}
