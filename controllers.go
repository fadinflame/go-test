package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateCustomerController(c *gin.Context) {
	var customer = Customer{Name: c.PostForm("name")}
	var created, err = storage.CreateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}
	c.JSON(http.StatusCreated, gin.H{"created": created.ID})
}

func GetCustomerController(c *gin.Context) {
	var customer, err = storage.getCustomerByName(c.PostForm("name"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Print(err.Error())
		return
	}
	c.JSON(http.StatusOK, customer)
}

func DepositMoneyController(c *gin.Context) {
	var amount, parseError = strconv.ParseFloat(c.PostForm("amount"), 64)
	if parseError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only floats available"})
		return
	}
	var res, err = storage.depositToBalance(c.PostForm("name"), amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !res {
		c.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong while depositing"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"amount": amount})
}

func IndexController(c *gin.Context) {
	var test string = "Albert"
	c.HTML(http.StatusOK, "index.html", gin.H{"test": test})
}
