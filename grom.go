package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Email string
}

func main() {
	//Connet to the mysql database
	dsn := "root:deadbyapril@15023@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("faild to connectet database")
	}
	//Autiomigrate the table
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("faild to auto migrate")
	}
	//Create a new user and some users
	singleUser := User{Name: "Sinan Chowdhury", Email: "sinanchowdhury23.com"}
	result := db.Create(&singleUser)
	if result.Error != nil {
		panic("Faild to create user ")
	}
	fmt.Println("Rows effected : ", result.RowsAffected)
	fmt.Println("Auto generated user ID : ", singleUser.ID)
	//CREATE MORE THAN ONE USER
	useArray := []User{
		{Name: "Dipu", Email: "dipu23@.com"},
		{Name: "Nipu", Email: "Nipu23@.com"},
		{Name: "Sipu", Email: "Sipu23@.com"},
		{Name: "Fipu", Email: "Fipu23@.com"},
	}
	for _, val := range useArray {

		result := db.Create(&val)
		if result.Error != nil {
			panic("Fail to create user")
		}

	}
	//Read a user
	var readUser User
	result = db.Where("email = ?", "dipu23@.com").First(&readUser)
	if result.Error != nil {
		panic("Faild to read user")
	}
	fmt.Printf("User: %v\n", readUser)
	//Read all user
	var readAllUser []User
	result = db.Find(&readAllUser)
	if result.Error != nil {
		panic("Faild to read user ")
	}
	fmt.Printf("All User: %v\n", readAllUser)
	fmt.Println("Raw effected : ", result.RowsAffected)
	//Update a user
	updateUser := User{ID: readUser.ID, Name: "Lipu", Email: "Lipu23@.com"}
	result = db.Save(&updateUser)
	if result.Error != nil {
		panic("Faild to updat")
	}
	// Delete a user
	result = db.Delete(&updateUser)
	if result.Error != nil {
		panic("Faild to delete user")
	}

}
