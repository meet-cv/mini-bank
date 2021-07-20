package services

import (
	"mini-bank/entities"
	"mini-bank/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func AccountInfo(db *pg.DB, c *gin.Context) {
	var result entities.Account
	acc_id := c.PostForm("acc_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("acc_id"), acc_id).Select()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.JSON(200, result)
	}
}

func AccountBalance(db *pg.DB, c *gin.Context) {
	var result entities.Account
	acc_id := c.PostForm("acc_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("acc_id"), acc_id).Select()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.JSON(200, result)
	}
}
