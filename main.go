package main

import (
	"database/sql"
	"log"
)

var storage *SQLiteRepository

//var customers []Customer

//func getCustomer(name string) (*Customer, error) {
//	var foundItem *Customer
//	for i, v := range customers {
//		if v.Name == name {
//			foundItem = &customers[i]
//			break
//		}
//	}
//	if foundItem == nil {
//		return foundItem, errors.New("could not find user by name")
//	}
//	return foundItem, nil
//}
//
//func createCustomerController(c *gin.Context) {
//	var customer = newCustomer(c.PostForm("name"))
//	customer.DepositCash(228)
//	customers = append(customers, customer)
//	c.JSON(http.StatusCreated, gin.H{"status": "created"})
//}

//func getCustomersController(c *gin.Context) {
//	c.JSON(http.StatusOK, customers)
//}

//func getCustomerBalanceController(c *gin.Context) {
//	var customer, err = getCustomer(c.Query("name"))
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"balance": customer.GetBalance()})
//}

//func setCustomerBalanceController(c *gin.Context) {
//	var customer, err = getCustomer(c.Query("name"))
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
//		return
//	}
//	amount, parseError := strconv.ParseFloat(c.PostForm("amount"), 64)
//	if parseError != nil {
//		c.JSON(http.StatusOK, gin.H{"error": "only floats available"})
//		return
//	}
//	customer.SetBalance(amount)
//	c.JSON(http.StatusOK, gin.H{customer.ID: "updated."})
//}


func main() {
	// database init
	const fileName = "./main.sqlite"
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	storage = NewSQLiteRepository(db)
	if err := storage.Migrate(); err != nil {
		log.Fatal(err)
	}
	r := NewRoutes()
	err = r.router.Run("localhost:8000")
	if err != nil {
		return
	}
}
