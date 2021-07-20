package services

import (
	"mini-bank/entities"
	"mini-bank/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

func AddUser(db *pg.DB, c *gin.Context) {
	var user entities.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		_, err := db.Model(&user).Insert()
		if err != nil {
			utils.ErrorResponse(c, err)
		} else {
			c.String(200, "user added")
		}
	}
}

func InfoUser(db *pg.DB, c *gin.Context) {
	var result entities.User
	user_id := c.PostForm("user_id")
	err := db.Model(&result).Where("? = ?", pg.Ident("user_id"), user_id).Select()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.JSON(200, result)
	}
}

func UpdateUser(db *pg.DB, c *gin.Context) { // IN PROGRESS
	var user entities.User
	c.BindJSON(&user)
	user.Updated_At = time.Now()
	_, err := db.Model(&user).ColumnExpr("").Insert()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.String(200, "user updated")
	}
}

func RemoveUser(db *pg.DB, c *gin.Context) {
	var user entities.User
	user_id := c.PostForm("user_id")
	_, err := db.Model(&user).Where("user_id = ?", user_id).Delete()
	if err != nil {
		utils.ErrorResponse(c, err)
	} else {
		c.String(200, "user deleted")
	}
}
