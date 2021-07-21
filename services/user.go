package services

import (
	"encoding/json"
	"fmt"
	"mini-bank/entities"
	"mini-bank/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func AddUser(db *pg.DB, c *gin.Context) {
	var user entities.UserAccount
	var res = make(map[string]string)
	err := c.ShouldBindJSON(&user)
	if !utils.ErrorResponse(c, err) {
		if !utils.ErrorResponse(c, err) {
			err = db.RunInTransaction(c, func(*pg.Tx) error {
				_, err = db.Model(&user.UserInfo).Insert()
				user.AccountInfo.Account_Number = fmt.Sprint(user.UserInfo.User_Id) + "AC"
				_, err = db.Model(&user.AccountInfo).Insert()
				var user_acc entities.AccountHolder
				user_acc.User_Id = user.UserInfo.User_Id
				user_acc.Account_Id = user.AccountInfo.Account_Id
				_, err = db.Model(&user_acc).Insert()
				res["message"] = fmt.Sprintf("user added with id : %d", user_acc.User_Id)
				return err
			})
			if !utils.ErrorResponse(c, err) {
				c.JSON(200, res)
			}
		}
	}
}

func InfoUser(db *pg.DB, c *gin.Context) {
	var result entities.User
	user_id := c.PostForm("user_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
	if err != nil && err.Error() != "pg: no rows in result set" {
		utils.ErrorResponse(c, err)
	} else if err != nil {
		var res = make(map[string]string)
		res["message"] = "no user found"
		c.JSON(400, res)
	} else {
		c.JSON(200, result)
	}
}

func UpdateUser(db *pg.DB, c *gin.Context) {
	raw_data, err := c.GetRawData()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		var result map[string]interface{}
		json.Unmarshal(raw_data, &result)
		_, err = db.Model(&result).TableExpr("users").Where("user_id = ?", result["user_id"]).Update()
		if err != nil {
			utils.ErrorResponse(c, err)
		} else {
			var res = make(map[string]string)
			res["message"] = fmt.Sprintf("user deleted with id : %s", result["user_id"])
			c.JSON(200, res)
		}
	}
}

func RemoveUser(db *pg.DB, c *gin.Context) {
	var user entities.User
	var res = make(map[string]string)
	user_id := c.PostForm("user_id")
	_, err := db.Model(&user).Where("user_id = ?", user_id).Delete()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		res["message"] = "user deleted"
		c.JSON(200, res)
	}
}
