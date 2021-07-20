package main

import (
	"log"
	"mini-bank/services"
	"mini-bank/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func main() {

	db = utils.DB_open()
	defer utils.DB_close(db)
	err := utils.DB_create(db)
	if err != nil {
		log.Printf("%v", err.Error())
	}

	router := gin.Default()
	user := router.Group("/user")
	{
		user.POST("/add", AddUser)
		user.POST("/info", InfoUser)
		user.POST("/update", UpdateUser)
		user.POST("/remove", RemoveUser)
		account := user.Group("/account")
		{
			account.POST("/info", AccountInfo)
			account.POST("/balance", AccountBalance)
		}
		transaction := user.Group("/transaction")
		{
			transaction.POST("/info", TransactionInfo)
			//transaction.POST("/execute", TrasactionExecute)
			transaction.POST("/all", TransactionAll)
		}
	}
	router.Run(":8080")
}

func AddUser(c *gin.Context) {
	services.AddUser(db, c)
}
func InfoUser(c *gin.Context) {
	services.InfoUser(db, c)
}
func UpdateUser(c *gin.Context) {
	services.UpdateUser(db, c)
}
func RemoveUser(c *gin.Context) {
	services.RemoveUser(db, c)
}
func AccountInfo(c *gin.Context) {
	services.AccountInfo(db, c)
}
func AccountBalance(c *gin.Context) {
	services.AccountBalance(db, c)
}
func TransactionInfo(c *gin.Context) {
	services.TransactionInfo(db, c)
}
func TransactionAll(c *gin.Context) {
	services.TransactionAll(db, c)
}
