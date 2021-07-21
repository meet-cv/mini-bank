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
	if err != nil && err.Error() != "pg: no rows in result set" {
		utils.ErrorResponse(c, err)
	} else if err != nil {
		var res = make(map[string]string)
		res["message"] = "no account found"
		c.JSON(400, res)
	} else {
		c.JSON(200, result)
	}
}

func AccountBalance(db *pg.DB, c *gin.Context) {
	var result entities.Account
	acc_id := c.PostForm("acc_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("acc_id"), acc_id).Select()
	if err != nil && err.Error() != "pg: no rows in result set" {
		utils.ErrorResponse(c, err)
	} else if err != nil {
		var res = make(map[string]string)
		res["message"] = "no account found"
		c.JSON(400, res)
	} else {
		var res = make(map[string]float64)
		res["balance"] = result.Account_Balance
		c.JSON(200, res)
	}
}
