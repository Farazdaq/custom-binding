package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomBind(c *gin.Context) {
	var request account

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"request": request,
	})
}

// THIS FUNCTION TO TAKE THE VALUES AND STORE THEM IN VARBLES
func BindAccount(c *gin.Context, accountnew *account) error {

	id := c.PostForm("ID")
	ac := c.PostForm("AcctiVation")
	full := c.PostForm("FullName")
	em := c.PostForm("Email")
	pass := c.PostForm("Password")
	acct := c.PostForm("AccountType")

}
