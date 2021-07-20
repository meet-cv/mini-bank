package services

import (
	"mini-bank/entities"
	"mini-bank/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func TransactionInfo(db *pg.DB, c *gin.Context) {
	var result entities.Transaction
	transac_id := c.PostForm("transac_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("transac_id"), transac_id).Select()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.JSON(200, result)
	}
}

func TransactionAll(db *pg.DB, c *gin.Context) {
	var result []entities.Transaction
	user_id := c.PostForm("user_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.JSON(200, result)
	}
}
