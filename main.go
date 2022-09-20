package main

import (
	"errors"
	"fmt"
	"learn-gorm/database"
	"learn-gorm/models"

	"gorm.io/gorm"
)

func main() {
	database.StartDB()

	// createUser("ally@mail.com")
	// getUserById(1)
	// createProduct(1, "YLO", "AA")
	getUsersWithProducts()
}

func createUser(email string) {
	db := database.GetDB()

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error
	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New user data:", User)
}

func getUserById(id uint) {
	db := database.GetDB()

	user := models.User{}
	err := db.First(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Println("User data:", user)
}

func createProduct(userId uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserId: userId,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}

	fmt.Println("New product data:", Product)
}

func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error

	if err != nil {
		fmt.Println("Error getting user data with products:", err.Error())
		return
	}

	fmt.Println("User data with products")
	fmt.Printf("%+v", users)
}
