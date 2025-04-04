package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	fmt.Println("migrated")

	var count int64
	db.Model(&User{}).Count(&count)
	if count == 0 {
		db.Create(&User{Name: "user01", Email: "xxxxxx@xxx01.com"})
		db.Create(&User{Name: "user02", Email: "xxxxxx@xxx02.com"})
		db.Create(&User{Name: "user03", Email: "xxxxxx@xxx03.com"})
	}

	var user User
	db.First(&user)
	fmt.Printf("変更前：%s",user.Name)

	db.Model(&user).Where("Email = ?", "xxxxxx@xxx01.com").Update("Name", "Yamada")
	db.First(&user)
	fmt.Printf("変更後：%s", user.Name)
}