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
	if err != nil && err.Error() != "pg: no rows in result set" {
		utils.ErrorResponse(c, err)
	} else if err != nil {
		var res = make(map[string]string)
		res["message"] = "no trasaction found"
		c.JSON(400, res)
	} else {
		c.JSON(200, result)
	}
}

func TransactionAll(db *pg.DB, c *gin.Context) {
	var result []entities.Transaction
	user_id := c.PostForm("user_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
	if err != nil && err.Error() != "pg: no rows in result set" {
		utils.ErrorResponse(c, err)
	} else if err != nil {
		var res = make(map[string]string)
		res["message"] = "no transaction found"
		c.JSON(400, res)
	} else {
		c.JSON(200, result)
	}
}

func TransactionExecute(db *pg.DB, c *gin.Context) {
	var transac entities.Transaction
	var res = make(map[string]interface{})
	err := c.ShouldBindJSON(&transac)
	if !utils.ErrorResponse(c, err) {
		var acc entities.Account
		var acchold entities.AccountHolder
		err = db.Model(&acchold).Where("user_id = ?", transac.User_Id).Select()
		err = db.Model(&acc).Where("acc_id = ?", acchold.Account_Id).Select()
		switch transac.Transaction_Type {
		case "credit":
			if transac.User_Id == transac.Beneficiary_Id {
				err = db.RunInTransaction(c, func(*pg.Tx) error {
					acc.Account_Balance += transac.Transaction_Amount
					_, err = db.Model(&acc).Where("acc_id = ?", acchold.Account_Id).Update()
					_, err = db.Model(&transac).Insert()
					err = db.Model(&acc).Where("acc_id = ?", acchold.Account_Id).Select()
					res["message"] = "account credited"
					res["transaction_id"] = transac.Transaction_Id
					res["account_info"] = acc
					return err
				})
				if !utils.ErrorResponse(c, err) {
					c.JSON(200, res)
				}
			}
		case "debit":
			if acc.Account_Balance < transac.Transaction_Amount {
				res["message"] = "insufficent account balance"
				c.JSON(400, res)
			} else {
				err = db.RunInTransaction(c, func(*pg.Tx) error {
					acc.Account_Balance -= transac.Transaction_Amount
					_, err = db.Model(&acc).Where("acc_id = ?", acchold.User_Id).Update()
					err = db.Model(&acchold).Where("user_id = ?", transac.Beneficiary_Id).Select()
					err = db.Model(&acc).Where("acc_id = ?", acchold.Account_Id).Select()
					acc.Account_Balance += transac.Transaction_Amount
					_, err = db.Model(&acc).Where("acc_id = ?", acchold.Account_Id).Update()
					_, err = db.Model(&transac).Insert()
					res["message"] = "account debited"
					res["transaction_id"] = transac.Transaction_Id
					return err
				})
				if !utils.ErrorResponse(c, err) {
					c.JSON(200, res)
				}
			}
		default:
			c.String(400, "unknown transaction type")
		}
	}
}
